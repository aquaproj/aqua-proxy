package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/suzuki-shunsuke/go-error-with-exit-code/ecerror"
)

type Runner struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

var errAquaCantBeExecuted = errors.New(`the command "aqua" can't be executed via aqua-proxy to prevent the infinite loop`)

func (runner *Runner) Run(ctx context.Context, args ...string) error {
	cmdName := filepath.Base(args[0])
	if cmdName == "aqua" {
		fmt.Fprintln(os.Stderr, "[ERROR] "+errAquaCantBeExecuted.Error())
		return errAquaCantBeExecuted
	}
	cmd := exec.CommandContext(ctx, "aqua", append([]string{"exec", "--", cmdName}, args[1:]...)...) //nolint:gosec
	cmd.Stdin = runner.Stdin
	cmd.Stdout = runner.Stdout
	cmd.Stderr = runner.Stderr

	setCancel(cmd)

	if err := cmd.Run(); err != nil {
		return ecerror.Wrap(err, cmd.ProcessState.ExitCode())
	}
	return nil
}

const waitDelay = 1000 * time.Hour

func setCancel(cmd *exec.Cmd) {
	cmd.Cancel = func() error {
		return cmd.Process.Signal(os.Interrupt)
	}
	cmd.WaitDelay = waitDelay
}

func absoluteAquaPath() (string, error) {
	aquaPath, err := exec.LookPath("aqua")
	if err != nil {
		return "", fmt.Errorf("aqua isn't found: %w", err)
	}
	if filepath.IsAbs(aquaPath) {
		return aquaPath, nil
	}
	a, err := filepath.Abs(aquaPath)
	if err != nil {
		return "", fmt.Errorf(`convert relative path "%s" to absolute path: %w`, aquaPath, err)
	}
	return a, nil
}
