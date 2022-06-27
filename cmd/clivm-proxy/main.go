package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/clivm/clivm-proxy/pkg/cli"
	"github.com/suzuki-shunsuke/go-error-with-exit-code/ecerror"
)

func main() {
	enabledXSysExec := os.Getenv("CLIVM_EXPERIMENTAL_X_SYS_EXEC") == "true" && runtime.GOOS != "windows"
	if err := core(enabledXSysExec); err != nil {
		if enabledXSysExec {
			fmt.Fprintln(os.Stderr, "[ERROR] "+err.Error())
			os.Exit(1)
		}
		os.Exit(ecerror.GetExitCode(err))
	}
}

func core(enabledXSysExec bool) error {
	runner := cli.Runner{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if enabledXSysExec {
		return runner.RunXSysExec(os.Args...) //nolint:wrapcheck
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	return runner.Run(ctx, os.Args...) //nolint:wrapcheck
}
