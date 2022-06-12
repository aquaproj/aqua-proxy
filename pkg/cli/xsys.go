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
