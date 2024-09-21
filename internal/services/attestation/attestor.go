package attestation

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/JITLiq/node/entity"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Attestor struct {
	pub         publisher
	keyManager  keyManager
	validator   validator
	blockReader ethereum.BlockNumberReader
}

func NewAttestor(
	pub publisher,
	manager keyManager,
	validator validator,
	blockReader ethereum.BlockNumberReader,
) *Attestor {
	return &Attestor{
		pub:         pub,
		keyManager:  manager,
		validator:   validator,
		blockReader: blockReader,
	}
}

func (a *Attestor) AttestAndPublish(ctx context.Context, msg entity.AttestOrderPayload) error {
	orderID, err := hexutil.Decode(msg.OrderID)
	if err != nil {
		return err
	}

	latest, err := a.blockReader.BlockNumber(ctx)
	if err != nil {
		return err
	}

	if err = a.validator.ValidateSolvedOrder(
		ctx,
		new(big.Int).SetInt64(int64(latest)),
		common.HexToHash(msg.OrderID), msg.Orders); err != nil {
		return err
	}

	sig, err := a.keyManager.Sign(ctx, hexutil.Encode(orderID))
	msg.Signature = hexutil.Encode(sig)
	msg.Operator = a.keyManager.Signer()
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return a.pub.Publish(ctx, data)
}

func (a *Attestor) PublishRequest(ctx context.Context, msg entity.AttestOrderPayload) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return a.pub.Publish(ctx, data)
}
