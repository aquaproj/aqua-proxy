//go:build !windows
// +build !windows

package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
)

func (runner *Runner) RunXSysExec(args ...string) error {
	cmdName := filepath.Base(args[0])
	if cmdName == "clivm" {
		return errAquaCantBeExecuted
	}

	clivmPath, err := absoluteAquaPath()
	if err != nil {
		return fmt.Errorf("get clivm's absolute path: %w", err)
	}
	if err := unix.Exec(clivmPath, append([]string{"clivm", "exec", "--", cmdName}, args[1:]...), os.Environ()); err != nil {
		return fmt.Errorf("execute clivm: %w", err)
	}
	return nil
}
