package solver

import (
	"context"
	"math/big"

	"github.com/JITLiq/node/entity"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type Solver struct {
	validator       srcStateManager
	manager         approvalManager
	destBlockReader ethereum.BlockNumberReader
	srcBlockReader  ethereum.BlockNumberReader
}

func NewSolver(
	validator srcStateManager,
	manager approvalManager,
	destBlockReader ethereum.BlockNumberReader,
	srcBlockReader ethereum.BlockNumberReader,
) (*Solver, error) {
	return &Solver{
		validator:       validator,
		manager:         manager,
		destBlockReader: destBlockReader,
		srcBlockReader:  srcBlockReader,
	}, nil
}

func (s *Solver) Solve(ctx context.Context, orderID common.Hash) (*entity.AttestOrderPayload, error) {
	latestSrc, err := s.srcBlockReader.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	order, err := s.validator.Order(ctx, new(big.Int).SetInt64(int64(latestSrc)), orderID)
	if err != nil {
		return nil, err
	}

	if err = s.validator.ValidatePendingOrder(ctx, new(big.Int).SetInt64(int64(latestSrc)), orderID); err != nil {
		return nil, err
	}

	latestDest, err := s.destBlockReader.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	solved, err := s.manager.Solve(ctx, new(big.Int).SetInt64(int64(latestDest)), order.OrderAmount)
	if err != nil {
		return nil, err
	}

	return &entity.AttestOrderPayload{
		OrderID: orderID.Hex(),
		Orders:  solved,
	}, nil
}
