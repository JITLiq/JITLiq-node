package attestation

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/JITLiq/node/entity"
	"github.com/JITLiq/node/pkg/logger"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"go.uber.org/zap"
)

const (
	Timeout    = time.Second * 20
	CheckEvery = time.Second * 2
)

type Aggregator struct {
	sub            subscriber
	hostID         string
	pendingOrderID string
	mu             sync.RWMutex
	sigs           []entity.AttestOrderPayload
	logger         *zap.Logger
	attestor       attestor
}

func NewAggregator(sub subscriber, attestor attestor) *Aggregator {
	return &Aggregator{
		sub:            sub,
		pendingOrderID: "",
		mu:             sync.RWMutex{},
		sigs:           make([]entity.AttestOrderPayload, 0),
		logger:         logger.NewLogger("aggregator"),
		attestor:       attestor,
	}
}

func (a *Aggregator) Aggregate(
	ctx context.Context,
	req entity.AttestOrderPayload,
	threshold int,
) ([]entity.AttestOrderPayload, error) {
	if a.pendingOrderID != "" {
		return nil, errors.New("pending aggregation")
	}
	a.mu.Lock()
	a.pendingOrderID = req.OrderID
	a.sigs = make([]entity.AttestOrderPayload, 0)
	a.mu.Unlock()

	defer func() {
		a.mu.Lock()
		a.pendingOrderID = req.OrderID
		a.sigs = make([]entity.AttestOrderPayload, 0)
		a.mu.Unlock()
	}()

	if err := a.attestor.PublishRequest(ctx, req); err != nil {
		return nil, err
	}

	ticker := time.Tick(CheckEvery)
	timeout := time.After(Timeout)
	for {
		select {
		case <-ticker:
			if len(a.sigs) == threshold {
				return a.sigs, nil
			}
		case <-timeout:
			if len(a.sigs) == threshold {
				return a.sigs, nil
			}

			return nil, errors.New("failed to aggregate sig, timeout occurred")
		}
	}
}

func (a *Aggregator) Run(ctx context.Context) error {
	sub, err := a.sub.Subscriber()
	if err != nil {
		return err
	}

	defer sub.Cancel()
	a.logger.Info("running attestation manager")
	for {
		fmt.Println("waiating for next")
		raw, err := sub.Next(ctx)
		if err != nil {
			return err
		}

		a.logger.Info("message received",
			zap.String("ID", string(raw.From)),
			zap.Any("topic", *raw.Topic),
			zap.Any("message", string(raw.Data)),
		)

		msg := &entity.AttestOrderPayload{}
		err = json.Unmarshal(raw.Data, msg)
		if err != nil {
			a.logger.Error("failed to parse", zap.Error(err))
			continue
		}

		// this is an attest request
		if msg.Signature == "" {
			fmt.Println("here")
			err = a.attestor.AttestAndPublish(ctx, *msg)
			if err != nil {
				a.logger.Error("failed to attest",
					zap.String("received", msg.OrderID),
					zap.Error(err))
			}
			continue
		}

		if a.pendingOrderID == "" {
			continue
		}

		if a.pendingOrderID != "" && msg.OrderID == a.pendingOrderID {
			a.mu.Lock()
			a.sigs = append(a.sigs, *msg)
			a.mu.Unlock()
		}

		if msg.OrderID != a.pendingOrderID {
			a.logger.Error("invalid orderID received",
				zap.String("pending", a.pendingOrderID),
				zap.String("received", msg.OrderID),
			)
		}

	}
}
