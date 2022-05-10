package main

import (
	"fmt"
	"os"

	"github.com/aquaproj/aqua-proxy/pkg/cli"
)

func main() {
	if err := core(); err != nil {
		fmt.Fprintln(os.Stderr, "[ERROR] "+err.Error())
		os.Exit(1)
	}
}

func core() error {
	runner := cli.Runner{
		Stderr: os.Stderr,
	}
	return runner.Run(os.Args...) //nolint:wrapcheck
}
