package repo

import (
	"context"
	"math/big"

	"github.com/JITLiq/node/abis/erc20"
	"github.com/JITLiq/node/entity"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

type ApprovalManager struct {
	state  entity.Fillers
	caller bind.ContractCaller
}

func NewApprovalManager(
	state entity.Fillers,
	caller bind.ContractCaller,
) *ApprovalManager {
	return &ApprovalManager{state: state, caller: caller}
}

func (a *ApprovalManager) Solve(ctx context.Context, at *big.Int, requestAmt *big.Int) ([]entity.Order, error) {
	caller, err := erc20.NewErc20Caller(entity.BaseUsdc, a.caller)
	if err != nil {
		return nil, err
	}

	pendingAmount := requestAmt
	orders := make([]entity.Order, 0)
	for _, filler := range a.state {
		bal, err := caller.BalanceOf(&bind.CallOpts{Context: ctx, BlockNumber: at}, filler)
		if err != nil {
			return nil, err
		}

		if bal.Cmp(new(big.Int).SetInt64(0)) == 0 {
			continue
		}

		allowance, err := caller.Allowance(&bind.CallOpts{Context: ctx, BlockNumber: at}, filler, entity.TargetJITPool)
		if err != nil {
			return nil, err
		}

		if pendingAmount.Cmp(new(big.Int).SetInt64(0)) == 0 {
			break
		}

		if bal.Cmp(requestAmt) > 0 && allowance.Cmp(pendingAmount) > 0 {
			return []entity.Order{
				{
					Filler: filler,
					Amount: entity.JsonBigInt{Int: *requestAmt},
				},
			}, nil
		}

		if bal.Cmp(allowance) >= 0 {
			orders = append(orders, entity.Order{
				Filler: filler,
				Amount: entity.JsonBigInt{Int: *bal},
			})

			pendingAmount.Sub(pendingAmount, bal)
		}
	}
	if pendingAmount.Cmp(new(big.Int).SetInt64(0)) != 0 {
		return nil, errors.New("not enough liquidity")
	}

	return orders, nil
}
