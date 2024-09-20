package attestation

import (
	"context"
	"encoding/json"

	"github.com/JITLiq/node/entity"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Attestor struct {
	pub        publisher
	keyManager keyManager
}

func NewAttestor(pub publisher, manager keyManager) *Attestor {
	return &Attestor{
		pub:        pub,
		keyManager: manager,
	}
}

func (a *Attestor) AttestAndPublish(ctx context.Context, msg entity.AttestOrderPayload) error {
	orderID, err := hexutil.Decode(msg.OrderID)
	if err != nil {
		return err
	}

	sig, err := a.keyManager.Sign(ctx, hexutil.Encode(crypto.Keccak256(orderID)))
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
