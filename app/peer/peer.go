package peer

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/JITLiq/node/config"
	"github.com/JITLiq/node/entity"
	"github.com/JITLiq/node/internal/integrations"
	"github.com/JITLiq/node/internal/repo"
	"github.com/JITLiq/node/internal/services/attestation"
	"github.com/JITLiq/node/internal/services/solver"
	"github.com/JITLiq/node/pkg/keymanager"
	"github.com/JITLiq/node/pkg/p2p"
	"github.com/ethereum/go-ethereum/common"
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

	signer, err := keymanager.NewKeyManager(cfg.Signer)
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

	srcStateManager, err := integrations.NewSrcStateManager(entity.SrcStateManager, arbClient)
	if err != nil {
		return err
	}

	aggSvc := attestation.NewAggregator(node, attestation.NewAttestor(node, signer, srcStateManager))

	solverSvc, err := solver.NewSolver(srcStateManager, approvalManager, baseClient, arbClient)
	if err != nil {
		panic(err)
	}

	if strings.Contains(configPath, "node1") {
		// go test(childCtx, aggSvc, approvalManager, srcStateManager, arbClient, baseClient)
		go processOrders(childCtx, cfg, solverSvc, aggSvc)
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

func test(
	ctx context.Context,
	agg *attestation.Aggregator,
	manager *repo.ApprovalManager,
	stateManager *integrations.SrcStateManager,
	src *ethclient.Client,
	dest *ethclient.Client,
) {
	<-time.After(time.Second * 5)

	solver, err := solver.NewSolver(stateManager, manager, dest, src)
	if err != nil {
		panic(err)
	}

	ord, err := solver.Solve(ctx, common.HexToHash("0xCFB31A4ECD2159118ED486050DDB966123D2FED1A0494A91EF255D3F2A156027"))
	if err != nil {
		panic(err)
	}

	aggSig, err := agg.Aggregate(ctx, *ord, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(aggSig)
}
