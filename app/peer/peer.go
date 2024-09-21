package peer

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/JITLiq/node/config"
	"github.com/JITLiq/node/entity"
	"github.com/JITLiq/node/internal/integrations"
	"github.com/JITLiq/node/internal/repo"
	"github.com/JITLiq/node/internal/services/attestation"
	"github.com/JITLiq/node/internal/services/relayer"
	"github.com/JITLiq/node/internal/services/solver"
	"github.com/JITLiq/node/pkg/keymanager"
	"github.com/JITLiq/node/pkg/p2p"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Run(ctx context.Context, configPath string) error {
	childCtx, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()

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

	if err = node.SetupPubSub(childCtx, entity.ORDER_PUBSUB_TOPIC); err != nil {
		return err
	}

	signer, err := keymanager.NewKeyManager(cfg.Signer, new(big.Int).SetInt64(entity.ChainIDBase))
	if err != nil {
		return err
	}

	baseClient, err := ethclient.Dial(cfg.BaseRPCURL)
	if err != nil {
		return err
	}

	arbClient, err := ethclient.Dial(cfg.ArbRPCURL)
	if err != nil {
		return err
	}

	approvalManager := repo.NewApprovalManager(entity.ToFillers(cfg.Fillers), baseClient)

	srcStateManager, err := integrations.NewSrcStateManager(entity.SrcStateManager, arbClient, baseClient)
	if err != nil {
		return err
	}

	aggSvc := attestation.NewAggregator(node, attestation.NewAttestor(node, signer, srcStateManager, baseClient))

	solverSvc, err := solver.NewSolver(srcStateManager, approvalManager, baseClient, arbClient)
	if err != nil {
		panic(err)
	}

	relayerClient, err := relayer.NewOrderRelayer(signer, baseClient)
	if err != nil {
		panic(err)
	}

	if strings.Contains(configPath, "node1") {
		// go test(childCtx, aggSvc, approvalManager, srcStateManager, arbClient, baseClient)
		go processOrders(childCtx, cfg, solverSvc, aggSvc, relayerClient)
	}

	go aggSvc.Run(childCtx)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("manager - Run - context canceled")
			return nil
		case s := <-interrupt:
			fmt.Println("manager - Run - signal: " + s.String())
			return nil
		}
	}
}
