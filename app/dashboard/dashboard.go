package dashboard

import (
	"context"
	"log"

	ui "github.com/gizak/termui/v3"
)

func Run(ctx context.Context, configPath string) error {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	cfg, err := LoadConfig(configPath)
	if err != nil {
		return err
	}

	components := createUIComponents()

	go updateOperatorsTable(ctx, cfg, components.lineChart, components.borrowPowerLog, components.operatorsTable)
	go updateLPTable(ctx, cfg, components.lineChartLP, components.totalLPValue)
	go updateOrdersTable(ctx, cfg, components.ordersTable)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		if e.ID == "q" || e.ID == "<C-c>" {
			return nil
		}
		renderAllComponents(components)
	}
}
