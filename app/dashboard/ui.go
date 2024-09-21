package dashboard

import (
	"context"
	"fmt"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type UIComponents struct {
	lineChart      *widgets.Plot
	lineChartLP    *widgets.Plot
	borrowPowerLog *widgets.Paragraph
	totalLPValue   *widgets.Paragraph
	ordersTable    *widgets.Table
	operatorsTable *widgets.Table
}

func createUIComponents() UIComponents {
	return UIComponents{
		lineChart:      createLineChart("Current JITLiq Borrow Power (USDC)", 0, 0, 50, 15),
		lineChartLP:    createLineChart("Current JITLiq Liquidity (USDC)", 52, 0, 100, 15),
		borrowPowerLog: createParagraph("Current JITLiq Borrow Power (USDC)", "0 USDC", 102, 0, 150, 6),
		totalLPValue:   createParagraph("Available JITLiq on Dest (USDC)", "0 USDC", 102, 8, 150, 14),
		ordersTable:    createTable("Orders", []string{"Order ID", "Transaction Hash", "Block"}, 0, 15, 70, 30),
		operatorsTable: createTable("Operators' Stake vs Holding", []string{
			"Operator",
			"Stake (USDC)",
			"Holding (USDC)",
		}, 72, 15, 150, 30),
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

func renderAllComponents(components UIComponents) {
	ui.Render(
		components.lineChart,
		components.borrowPowerLog,
		components.totalLPValue,
		components.ordersTable,
		components.operatorsTable,
		components.lineChartLP,
	)
}
