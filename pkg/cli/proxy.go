package cli

import (
	"context"
	"fmt"
	"io"
	"os/exec"

	"github.com/suzuki-shunsuke/go-timeout/timeout"
)

type Runner struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func (runner *Runner) Run(ctx context.Context, args ...string) error {
	cmd := exec.Command("aqua", append([]string{"exec", "--"}, args...)...) //nolint:gosec
	cmd.Stdin = runner.Stdin
	cmd.Stdout = runner.Stdout
	cmd.Stderr = runner.Stderr
	r := timeout.NewRunner(0)
	if err := r.Run(ctx, cmd); err != nil {
		return fmt.Errorf("execute the command: %w", err)
	}
	return nil
}
