package dashboard

import (
	"context"
	"math/big"

	"github.com/JITLiq/node/abis/erc20"
	"github.com/JITLiq/node/abis/srcstatemanager"
	"github.com/JITLiq/node/entity"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CurrentHoldings(ctx context.Context, cfg *Config) ([]OperatorsStake, error) {
	client, err := ethclient.Dial(cfg.ArbRPCURL)
	if err != nil {
		return nil, err
	}

	manager, err := srcstatemanager.NewSrcstatemanager(entity.SrcStateManager, client)
	if err != nil {
		return nil, err
	}

	stake := make([]OperatorsStake, 0, len(cfg.Operators))
	for _, op := range cfg.Operators {
		data, err := manager.OperatorData(&bind.CallOpts{Context: ctx}, common.HexToAddress(op))
		if err != nil {
			return nil, err
		}

		stake = append(stake, OperatorsStake{
			Holding: float64(data.CurrentHolding.Uint64()) / 1e6,
			Staked:  float64(data.CurrentStake.Uint64()) / 1e6,
			Address: op,
		})
	}

	return stake, nil
}

func CurrentLP(ctx context.Context, cfg *Config) (float64, error) {
	client, err := ethclient.Dial(cfg.BaseRPCURL)
	if err != nil {
		return 0, err
	}

	tokenCaller, err := erc20.NewErc20Caller(entity.BaseUsdc, client)
	if err != nil {
		return 0, err
	}

	var available float64
	for _, filler := range cfg.Fillers {
		data, err := tokenCaller.BalanceOf(&bind.CallOpts{Context: ctx}, common.HexToAddress(filler))
		if err != nil {
			return 0, err
		}
		available += float64(data.Uint64()) / 1e6
	}

	return available, nil
}

func FetchOrders(ctx context.Context, cfg *Config) error {
	client, err := ethclient.Dial(cfg.ArbRPCURL)
	if err != nil {
		return err
	}

	current, err := client.BlockNumber(ctx)
	if err != nil {
		return err
	}

	logs, err := client.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: new(big.Int).SetInt64(int64(current) - 200),
		ToBlock:   new(big.Int).SetInt64(int64(current)),
		Addresses: []common.Address{entity.SrcStateManager},
		Topics:    [][]common.Hash{{entity.TopicOrderCreated}},
	})
	if err != nil {
		return err
	}

	for _, l := range logs {
		if len(l.Topics) < 2 {
			continue
		}
		orderID := l.Topics[1].Hex()
		if _, ok := orders[orderID]; ok {
			continue
		}
		orders[orderID] = Order{
			TxHash:   l.TxHash.Hex(),
			OrderID:  orderID,
			BlockNum: l.BlockNumber,
		}
	}
	return nil
}
