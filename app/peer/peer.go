package peer

import (
	"context"
	"fmt"
	"strings"
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

	baseClient, err := ethclient.Dial(cfg.BaseRPCURL)
	if err != nil {
		return err
	}

	arbClient, err := ethclient.Dial(cfg.ArbRPCURL)
	if err != nil {
		return err
	}

	approvalManager := repo.NewApprovalManager(entity.ToFillers(cfg.Fillers), baseClient)

	srcStateManager, err := integrations.NewSrcStateManager(common.HexToAddress("0x2bCFbC4Dd2Af9Af0458c9aD179A7A5791b0A9502"), arbClient)
	if err != nil {
		return err
	}

	svc := attestation.NewAggregator(node, attestation.NewAttestor(node, signer, srcStateManager))
	if strings.Contains(configPath, "node1") {
		go test(ctx, svc, approvalManager, srcStateManager, arbClient, baseClient)
	}
	return svc.Run(ctx)
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

	solvver, err := solver.NewSolver(stateManager, manager, dest, src)
	if err != nil {
		panic(err)
	}

	ord, err := solvver.Solve(ctx, common.HexToHash(""))
	aggSig, err := agg.Aggregate(ctx, *ord, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(aggSig)
}
