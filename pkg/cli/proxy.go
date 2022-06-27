package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/suzuki-shunsuke/go-error-with-exit-code/ecerror"
	"github.com/suzuki-shunsuke/go-timeout/timeout"
)

type Runner struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

var errAquaCantBeExecuted = errors.New(`the command "clivm" can't be executed via clivm-proxy to prevent the infinite loop`)

func (runner *Runner) Run(ctx context.Context, args ...string) error {
	cmdName := filepath.Base(args[0])
	if cmdName == "clivm" {
		fmt.Fprintln(os.Stderr, "[ERROR] "+errAquaCantBeExecuted.Error())
		return errAquaCantBeExecuted
	}
	cmd := exec.Command("clivm", append([]string{"exec", "--", cmdName}, args[1:]...)...) //nolint:gosec
	cmd.Stdin = runner.Stdin
	cmd.Stdout = runner.Stdout
	cmd.Stderr = runner.Stderr
	r := timeout.NewRunner(0)
	if err := r.Run(ctx, cmd); err != nil {
		return ecerror.Wrap(err, cmd.ProcessState.ExitCode())
	}
	return nil
}

func absoluteAquaPath() (string, error) {
	clivmPath, err := exec.LookPath("clivm")
	if err != nil {
		return "", fmt.Errorf("clivm isn't found: %w", err)
	}
	if filepath.IsAbs(clivmPath) {
		return clivmPath, nil
	}
	a, err := filepath.Abs(clivmPath)
	if err != nil {
		return "", fmt.Errorf(`convert relative path "%s" to absolute path: %w`, clivmPath, err)
	}
	return a, nil
}
