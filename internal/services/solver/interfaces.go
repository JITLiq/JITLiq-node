package solver

import (
	"context"
	"math/big"

	"github.com/JITLiq/node/abis/srcstatemanager"
	"github.com/JITLiq/node/entity"
	"github.com/ethereum/go-ethereum/common"
)

type approvalManager interface {
	Solve(ctx context.Context, at *big.Int, requestAmt *big.Int) ([]entity.Order, error)
}

type srcStateManager interface {
	ValidateSolvedOrder(
		ctx context.Context,
		at *big.Int,
		orderID common.Hash,
		orders []entity.Order,
	) error
	ValidatePendingOrder(
		ctx context.Context,
		at *big.Int,
		orderID common.Hash,
	) error
	Order(ctx context.Context, at *big.Int, orderID common.Hash) (struct {
		Fulfilled   bool
		Expiry      uint32
		OrderAmount *big.Int
		DestAddress common.Address
		Operator    common.Address
		Fees        srcstatemanager.ISourceOpStateManagerFeesData
	}, error)
}
