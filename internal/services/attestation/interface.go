package attestation

import (
	"context"

	"github.com/JITLiq/node/entity"
	"github.com/ethereum/go-ethereum/common"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

type publisher interface {
	Publish(ctx context.Context, msg []byte) error
}

type subscriber interface {
	Subscriber() (*pubsub.Subscription, error)
}

type keyManager interface {
	Sign(_ context.Context, hash string) ([]byte, error)
	Signer() common.Address
}

type attestor interface {
	AttestAndPublish(ctx context.Context, msg entity.AttestOrderPayload) error
	PublishRequest(ctx context.Context, msg entity.AttestOrderPayload) error
}
