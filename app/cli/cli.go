package cli

import (
	"context"

	"github.com/JITLiq/node/app/peer"
	"github.com/urfave/cli/v3"
)

var (
	envPath = ""
)

func BuildCLI() *cli.Command {
	return &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "node",
				Aliases: []string{"node"},
				Usage:   "Runs http server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "env",
						Required:    true,
						Destination: &envPath,
					},
				},
				Action: func(ctx context.Context, _ *cli.Command) error {
					return peer.Run(ctx, envPath)
				},
			},
		},
	}
}
