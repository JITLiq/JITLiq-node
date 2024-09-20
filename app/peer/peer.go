package peer

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/JITLiq/node/config"
	"github.com/JITLiq/node/entity"
	"github.com/JITLiq/node/internal/services/attestation"
	"github.com/JITLiq/node/pkg/keymanager"
	"github.com/JITLiq/node/pkg/p2p"
	"github.com/ethereum/go-ethereum/common"
)

func Run(ctx context.Context, configPath string) error {
	cfg := &config.Config{}
	err := config.LoadEnvConfig(cfg, config.WithEnvPath(configPath))
	if err != nil {
		return err
	}

	node, err := p2p.New(&p2p.NodeConfig{Addr: cfg.P2PHost})
	if err != nil {
		return err
	}

	if err = node.SetupDiscovery(); err != nil {
		return err
	}

	if err = node.SetupPubSub(ctx, entity.ORDER_PUBSUB_TOPIC); err != nil {
		return err
	}

	signer, err := keymanager.NewKeyManager(cfg.Signer)
	if err != nil {
		return err
	}

	svc := attestation.NewAggregator(node, attestation.NewAttestor(node, signer))
	if strings.Contains(configPath, "node1") {
		go test(ctx, svc)
	}
	return svc.Run(ctx)
}

func test(ctx context.Context, agg *attestation.Aggregator) {
	<-time.After(time.Second * 5)

	req := entity.AttestOrderPayload{
		OrderID:         "0xabcd",
		TransactionHash: common.HexToHash("0x1234"),
	}

	resp, err := agg.Aggregate(ctx, req, 3)
	fmt.Println(resp, err)
}
