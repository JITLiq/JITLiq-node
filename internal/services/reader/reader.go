package reader

import (
	"context"
	"math/big"
	"time"

	"github.com/JITLiq/node/entity"
	"github.com/JITLiq/node/pkg/logger"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

const (
	PollingInterval = time.Second * 15
	LogBuffer       = 1
)

type OrderReader struct {
	lastCheckedBlock  int64
	filterer          ethereum.LogFilterer
	blockNumberReader ethereum.BlockNumberReader
	sink              chan types.Log
}

func NewOrderReader(
	initBlock int64,
	filterer ethereum.LogFilterer,
	blockNumberReader ethereum.BlockNumberReader,
) *OrderReader {
	return &OrderReader{
		lastCheckedBlock:  initBlock,
		filterer:          filterer,
		sink:              make(chan types.Log, LogBuffer),
		blockNumberReader: blockNumberReader,
	}
}

func (o *OrderReader) Read() <-chan types.Log {
	return o.sink
}

func (o *OrderReader) Subscribe(ctx context.Context) {
	log := logger.NewLogger("OrderReader")
	poller := time.Tick(PollingInterval)
	for {
		select {
		case <-ctx.Done():
			return
		case <-poller:
			latest, err := o.blockNumberReader.BlockNumber(ctx)
			if err != nil {
				log.Error("failed to read latest block", zap.Error(err))
				continue
			}

			if int64(latest)-o.lastCheckedBlock > 2000 {
				latest += uint64(o.lastCheckedBlock + 2000)
			}

			logs, err := o.filterer.FilterLogs(ctx, ethereum.FilterQuery{
				FromBlock: new(big.Int).SetInt64(o.lastCheckedBlock),
				ToBlock:   new(big.Int).SetInt64(int64(latest)),
				Topics: [][]common.Hash{
					{
						entity.TopicOrderCreated,
					},
				},
			})
			if err != nil {
				log.Error("failed to query orders", zap.Error(err))
				continue
			}

			log.Info("fetched blocks",
				zap.Int("from", int(o.lastCheckedBlock)),
				zap.Int("latest", int(latest)),
				zap.Any("logs", logs),
			)

			o.lastCheckedBlock = int64(latest)

			if len(logs) == 0 {
				continue
			}

			for _, l := range logs {
				o.sink <- l
			}

		}

	}
}
