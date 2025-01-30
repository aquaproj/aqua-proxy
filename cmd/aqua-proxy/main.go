package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/aquaproj/aqua-proxy/pkg/cli"
	"github.com/suzuki-shunsuke/go-error-with-exit-code/ecerror"
)

func main() {
	enabledXSysExec := getEnabledXSysExec(runtime.GOOS)
	if err := core(enabledXSysExec); err != nil {
		if enabledXSysExec {
			fmt.Fprintln(os.Stderr, "[ERROR] "+err.Error())
			os.Exit(1)
		}
		os.Exit(ecerror.GetExitCode(err))
	}
}

func getEnabledXSysExec(goos string) bool {
	if goos == "windows" {
		return false
	}
	if os.Getenv("AQUA_EXPERIMENTAL_X_SYS_EXEC") == "false" {
		return false
	}
	if os.Getenv("AQUA_X_SYS_EXEC") == "false" {
		return false
	}
	return true
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
