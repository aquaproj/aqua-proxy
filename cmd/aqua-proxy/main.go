package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/suzuki-shunsuke/aqua-proxy/pkg/cli"
)

func main() {
	if err := core(); err != nil {
		log.Fatal(err)
	}
}

func core() error {
	runner := cli.Runner{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	return runner.Run(ctx, os.Args...) //nolint:wrapcheck
}
