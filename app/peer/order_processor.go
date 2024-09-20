package peer

import (
	"context"
	"fmt"

	"github.com/JITLiq/node/config"
	"github.com/JITLiq/node/internal/services/attestation"
	"github.com/JITLiq/node/internal/services/reader"
	"github.com/JITLiq/node/internal/services/solver"
	"github.com/JITLiq/node/pkg/logger"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"go.uber.org/zap"
)

const (
	minQuorum = 2
)

func processOrders(
	ctx context.Context,
	cfg *config.Config,
	solver *solver.Solver,
	aggregator *attestation.Aggregator,
) {
	log := logger.NewLogger("OrderProcessor")
	client, err := ethclient.Dial(cfg.ArbRPCURL)
	if err != nil {
		log.Error("failed to dail ethclient", zap.Error(err))
		return
	}
	latest, err := client.BlockNumber(ctx)
	if err != nil {
		log.Error("failed to fetch latest block", zap.Error(err))
	}

	sub := reader.NewOrderReader(int64(latest), client, client)
	go sub.Subscribe(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		case l := <-sub.Read():
			err = solveAndPrepareForRelay(ctx, l, solver, aggregator)
			if err != nil {
				log.Error("failed to solve order", zap.Error(err))
			}
		}
	}
}

func solveAndPrepareForRelay(
	ctx context.Context,
	log types.Log,
	solver *solver.Solver,
	aggregator *attestation.Aggregator,
) error {
	if len(log.Topics) < 2 {
		return errors.New("topics too smol")
	}
	orderID := log.Topics[1]
	payload, err := solver.Solve(ctx, orderID)
	if err != nil {
		return err
	}

	attestations, err := aggregator.Aggregate(ctx, *payload, minQuorum)
	if err != nil {
		return err
	}

	fmt.Println(attestations)
	return nil
}
