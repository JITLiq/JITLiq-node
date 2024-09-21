package dashboard

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/JITLiq/node/abis/erc20"
	"github.com/JITLiq/node/abis/srcstatemanager"
	"github.com/JITLiq/node/config"
	"github.com/JITLiq/node/entity"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const UpdateInterval = 5 * time.Second

type Config struct {
	BaseRPCURL string   `envconfig:"BASE_RPC_URL"`
	ArbRPCURL  string   `envconfig:"ARB_RPC_URL"`
	Fillers    []string `envconfig:"FILLERS"`
	Operators  []string `envconfig:"OPERATORS"`
}

type OperatorsStake struct {
	Holding float64
	Staked  float64
	Address string
}

type Order struct {
	TxHash   string
	OrderID  string
	BlockNum uint64
}

var orders = map[string]Order{}

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

func Run(ctx context.Context, configPath string) error {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	cfg := &Config{}
	if err := config.LoadEnvConfig(cfg, config.WithEnvPath(configPath)); err != nil {
		return err
	}

	lineChart := createLineChart("Current JITLiq Borrow Power (USDC)", 0, 0, 50, 15)
	lineChartLP := createLineChart("Current JITLiq Liquidity (USDC)", 52, 0, 100, 15)
	borrowPowerLog := createParagraph("Current JITLiq Borrow Power (USDC)", "0 USDC", 102, 0, 150, 6)
	totalLPValue := createParagraph("Available JITLiq on Dest (USDC)", "0 USDC", 102, 8, 150, 14)
	ordersTable := createTable("Orders", []string{"Order ID", "Transaction Hash", "Block"}, 0, 15, 70, 30)
	operatorsTable := createTable("Operators' Stake vs Holding", []string{
		"Operator",
		"Stake (USDC)",
		"Holding (USDC)",
	}, 72, 15, 150, 30)

	go updateOperatorsTable(ctx, cfg, lineChart, borrowPowerLog, operatorsTable)
	go updateLPTable(ctx, cfg, lineChartLP, totalLPValue)
	go updateOrdersTable(ctx, cfg, ordersTable)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		if e.ID == "q" || e.ID == "<C-c>" {
			return nil
		}
		ui.Render(lineChart, borrowPowerLog, totalLPValue, ordersTable, operatorsTable, lineChartLP)
	}
}

func createLineChart(title string, x1, y1, x2, y2 int) *widgets.Plot {
	chart := widgets.NewPlot()
	chart.Title = title
	chart.Data = make([][]float64, 3)
	for i := range chart.Data {
		chart.Data[i] = make([]float64, 10)
	}
	chart.SetRect(x1, y1, x2, y2)
	chart.AxesColor = ui.ColorWhite
	chart.LineColors[0] = ui.ColorMagenta
	chart.LineColors[1] = ui.ColorMagenta
	chart.LineColors[2] = ui.ColorMagenta
	return chart
}

func createParagraph(title, text string, x1, y1, x2, y2 int) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Title = title
	p.Text = text
	p.SetRect(x1, y1, x2, y2)
	p.TextStyle = ui.NewStyle(ui.ColorWhite, ui.ColorClear, ui.ModifierBold)
	return p
}

func createTable(title string, headers []string, x1, y1, x2, y2 int) *widgets.Table {
	table := widgets.NewTable()
	table.Title = title
	table.Rows = [][]string{headers}
	table.SetRect(x1, y1, x2, y2)
	table.TextStyle = ui.NewStyle(ui.ColorWhite, ui.ColorClear, ui.ModifierBold)
	table.RowSeparator = true
	table.BorderStyle.Fg = ui.ColorYellow
	return table
}

func updateOperatorsTable(
	ctx context.Context,
	cfg *Config,
	lineChart *widgets.Plot,
	borrowPowerLog *widgets.Paragraph,
	operatorsTable *widgets.Table,
) {
	for {
		holdings, err := CurrentHoldings(ctx, cfg)
		if err != nil {
			time.Sleep(UpdateInterval)
			continue
		}

		var totalStaked, totalBorrowed float64
		operatorsTable.Rows = [][]string{{"Operator", "Stake (USDC)", "Holding (USDC)"}}

		for _, h := range holdings {
			operatorsTable.Rows = append(operatorsTable.Rows, []string{
				h.Address,
				fmt.Sprintf("%.4f", h.Staked),
				fmt.Sprintf("%.4f", h.Holding),
			})
			totalBorrowed += h.Holding
			totalStaked += h.Staked
		}

		borrowPower := totalStaked - totalBorrowed
		for i := range lineChart.Data {
			lineChart.Data[i] = append(lineChart.Data[i][1:], borrowPower)
		}
		borrowPowerLog.Text = fmt.Sprintf("%.2f USDC", borrowPower)

		ui.Render(lineChart, borrowPowerLog, operatorsTable)
		time.Sleep(UpdateInterval)
	}
}

func updateLPTable(ctx context.Context, cfg *Config, lineChartLP *widgets.Plot, totalLPValue *widgets.Paragraph) {
	for {
		lp, err := CurrentLP(ctx, cfg)
		if err != nil {
			time.Sleep(UpdateInterval)
			continue
		}

		totalLPValue.Text = fmt.Sprintf("%.4f USDC", lp)
		for i := range lineChartLP.Data {
			lineChartLP.Data[i] = append(lineChartLP.Data[i][1:], lp)
		}

		ui.Render(lineChartLP, totalLPValue)
		time.Sleep(UpdateInterval)
	}
}

func updateOrdersTable(ctx context.Context, cfg *Config, ordersTable *widgets.Table) {
	for {
		if err := FetchOrders(ctx, cfg); err != nil {
			time.Sleep(UpdateInterval)
			continue
		}

		ordersTable.Rows = [][]string{{"Order ID", "Transaction Hash", "Block"}}
		for _, o := range orders {
			ordersTable.Rows = append(ordersTable.Rows, []string{
				o.OrderID,
				o.TxHash,
				fmt.Sprintf("%d", o.BlockNum),
			})
		}

		ui.Render(ordersTable)
		time.Sleep(UpdateInterval)
	}
}
