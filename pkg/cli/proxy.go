package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/sys/unix"
)

type Runner struct {
	Stderr io.Writer
}

var errAquaCantBeExecuted = errors.New(`the command "aqua" can't be executed via aqua-proxy to prevent the infinite loop`)

func (runner *Runner) Run(args ...string) error {
	cmdName := filepath.Base(args[0])
	if cmdName == "aqua" {
		return errAquaCantBeExecuted
	}

	aquaPath, err := absoluteAquaPath()
	if err != nil {
		return fmt.Errorf("get aqua's absolute path: %w", err)
	}
	if err := unix.Exec(aquaPath, append([]string{"aqua", "exec", "--", cmdName}, args[1:]...), os.Environ()); err != nil {
		return fmt.Errorf("execute aqua: %w", err)
	}
	return nil
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
