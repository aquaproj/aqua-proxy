package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/suzuki-shunsuke/aqua-proxy/pkg/cli"
	"github.com/suzuki-shunsuke/go-error-with-exit-code/ecerror"
)

func main() {
	if err := core(); err != nil {
		os.Exit(ecerror.GetExitCode(err))
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
