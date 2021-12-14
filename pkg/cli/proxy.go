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

var errAquaCantBeExecuted = errors.New(`the command "aqua" can't be executed via aqua-proxy to prevent the infinite loop`)

func (runner *Runner) Run(ctx context.Context, args ...string) error {
	cmdName := filepath.Base(args[0])
	if cmdName == "aqua" {
		fmt.Fprintln(os.Stderr, "[ERROR] "+errAquaCantBeExecuted.Error())
		return errAquaCantBeExecuted
	}
	cmd := exec.Command("aqua", append([]string{"exec", "--", cmdName}, args[1:]...)...) //nolint:gosec
	cmd.Stdin = runner.Stdin
	cmd.Stdout = runner.Stdout
	cmd.Stderr = runner.Stderr
	r := timeout.NewRunner(0)
	if err := r.Run(ctx, cmd); err != nil {
		return ecerror.Wrap(err, cmd.ProcessState.ExitCode())
	}
	return nil
}
