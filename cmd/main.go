package main

import (
	"context"
	"log"
	"os"

	"github.com/JITLiq/node/app/cli"
)

func main() {
	err := cli.BuildCLI().Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
