package relayer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/JITLiq/node/abis/destentrypoint"
	"github.com/JITLiq/node/abis/deststatemanager"
	"github.com/JITLiq/node/entity"
	"github.com/JITLiq/node/pkg/logger"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"go.uber.org/zap"
)

const (
	MaxReceiptFetch = 10
	SleepTime       = time.Second * 4
)

type OrderRelayer struct {
	signer     keymanager
	transactor *destentrypoint.Destentrypoint
	reader     *deststatemanager.DeststatemanagerCaller
	fetcher    ethereum.TransactionReader
}

func NewOrderRelayer(signer keymanager, client *ethclient.Client) (*OrderRelayer, error) {
	transactor, err := destentrypoint.NewDestentrypoint(entity.TargetJITPool, client)
	if err != nil {
		return nil, err
	}

	reader, err := deststatemanager.NewDeststatemanagerCaller(entity.TargetStateManager, client)
	if err != nil {
		return nil, err
	}

	return &OrderRelayer{
		transactor: transactor,
		reader:     reader,
		fetcher:    client,
		signer:     signer,
	}, nil
}

func (o *OrderRelayer) RelayOrder(ctx context.Context, payloads []entity.AttestOrderPayload) (*common.Hash, error) {
	log := logger.NewLogger("OrderRelayer")
	operators, err := o.reader.GetOperators(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	signatures := make(map[common.Address]string)
	for _, p := range payloads {
		signatures[p.Operator] = p.Signature
	}

	sigBytes := make([][]byte, 0)
	for _, op := range operators {
		sigStr, ok := signatures[op]
		if !ok {
			return nil, fmt.Errorf("failed to get signature %s", op.String())
		}

		sig, err := hexutil.Decode(sigStr)
		if err != nil {
			return nil, err
		}

		sigBytes = append(sigBytes, sig)
	}

	payload := payloads[0]
	orderID := common.HexToHash(payload.OrderID)
	orderData := destentrypoint.IEntityOrderData{
		Fulfilled:   payload.Data.Fulfilled,
		Expiry:      payload.Data.Expiry,
		OrderAmount: payload.Data.OrderAmount,
		DestAddress: payload.Data.DestAddress,
		Operator:    payload.Data.Operator,
		Fees: destentrypoint.IEntityFeesData{
			OperationFee: payload.Data.Fees.OperationFee,
			BridgeFee:    payload.Data.Fees.BridgeFee,
		},
	}

	fillers := make([]destentrypoint.IEntityFulfillerData, 0)
	for _, p := range payload.Orders {
		fillers = append(fillers, destentrypoint.IEntityFulfillerData{
			FulfillAmount: &p.Amount.Int,
			Fulfiller:     p.Filler,
		})
	}
	log.Info("relaying payload")

	order, err := o.transactor.FulfillOrder(&bind.TransactOpts{
		Signer:  o.signer.TransactionSignerHook,
		Context: ctx,
		From:    o.signer.Signer(),
		Value:   new(big.Int).SetInt64(28237182426157),
	}, orderID, orderData, fillers, sigBytes)
	if err != nil {
		return nil, err
	}

	log.Info("generated tx data", zap.Any("calldata", hexutil.Encode(order.Data())), zap.Any("target", order.To()))
	log.Info("generated tx", zap.Any("order", order.Hash()))

	for i := 0; i < MaxReceiptFetch; i++ {
		receipt, err := o.fetcher.TransactionReceipt(ctx, order.Hash())
		if err != nil {
			log.Error("failed to fetch receipt", zap.Error(err))
			<-time.After(SleepTime)
			continue
		}
		if receipt == nil {
			log.Error("receipt not found")
			<-time.After(SleepTime)
			continue
		}

		log.Info("receipt fetched, order relayed", zap.Any("minedTx", receipt.TxHash))
		return &receipt.TxHash, nil
	}

	return nil, errors.New("failed to fetch receipt")

}
