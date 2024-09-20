package integrations

import (
	"context"
	"fmt"
	"math/big"

	"github.com/JITLiq/node/abis/erc20"
	"github.com/JITLiq/node/abis/srcstatemanager"
	"github.com/JITLiq/node/entity"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

type SrcStateManager struct {
	managerCaller  *srcstatemanager.SrcstatemanagerCaller
	contractReader bind.ContractCaller
}

func NewSrcStateManager(addr common.Address, caller bind.ContractCaller) (*SrcStateManager, error) {
	c, err := srcstatemanager.NewSrcstatemanagerCaller(addr, caller)
	if err != nil {
		return nil, err
	}

	return &SrcStateManager{managerCaller: c, contractReader: caller}, nil
}
func (s *SrcStateManager) ValidatePendingOrder(
	ctx context.Context,
	at *big.Int,
	orderID common.Hash,
) error {
	orderData, err := s.managerCaller.OrderData(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: at,
	}, orderID)

	if err != nil {
		return err
	}

	if !isPendingOrder(at, orderData.Expiry, orderData.Fulfilled) {
		return errors.New("order already fulfilled or expired")
	}

	return nil
}

func (s *SrcStateManager) Order(ctx context.Context, at *big.Int, orderID common.Hash) (struct {
	Fulfilled   bool
	Expiry      uint32
	OrderAmount *big.Int
	DestAddress common.Address
	Operator    common.Address
	Fees        srcstatemanager.ISourceOpStateManagerFeesData
}, error) {
	orderData, err := s.managerCaller.OrderData(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: at,
	}, orderID)

	if err != nil {
		return struct {
			Fulfilled   bool
			Expiry      uint32
			OrderAmount *big.Int
			DestAddress common.Address
			Operator    common.Address
			Fees        srcstatemanager.ISourceOpStateManagerFeesData
		}{}, nil
	}

	return orderData, nil
}
func (s *SrcStateManager) ValidateSolvedOrder(
	ctx context.Context,
	at *big.Int,
	orderID common.Hash,
	orders []entity.Order,
) error {
	orderData, err := s.managerCaller.OrderData(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: at,
	}, orderID)

	if err != nil {
		return err
	}

	if err = s.validateFillers(ctx, entity.TargetJITPool, orders, orderData.OrderAmount); err != nil {
		return err
	}

	return nil
}

func (s *SrcStateManager) validateFillers(
	ctx context.Context,
	target common.Address,
	orders []entity.Order,
	requestedAmt *big.Int,
) error {
	pulledAmt := new(big.Int).SetInt64(0)
	erc20Reader, err := erc20.NewErc20Caller(entity.BaseUsdc, s.contractReader)
	if err != nil {
		return err
	}

	for _, o := range orders {
		allowance, err := erc20Reader.Allowance(&bind.CallOpts{Context: ctx}, o.Filler, target)
		if err != nil {
			return err
		}

		if allowance.Cmp(&o.Amount.Int) < 0 {
			return errors.New("allowance too low")
		}

		bal, err := erc20Reader.BalanceOf(&bind.CallOpts{
			Context: ctx,
		}, o.Filler)
		if err != nil {
			return err
		}

		if bal.Cmp(&o.Amount.Int) < 0 {
			return errors.New("not enough balance")
		}

		pulledAmt.Add(pulledAmt, &o.Amount.Int)
	}

	if pulledAmt.Cmp(requestedAmt) != 0 {
		return fmt.Errorf("order mismatch, pulled: %s requested %s", pulledAmt.String(), requestedAmt.String())
	}

	return nil
}

func isPendingOrder(at *big.Int, expiry uint32, fulfilled bool) bool {
	// TODO: add block check
	//  && uint64(expiry) > at.Uint64()
	return !fulfilled
}
