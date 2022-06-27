//go:build windows
// +build windows

package cli

import "errors"

var errXSysNotSuppported = errors.New("Windows doesn't support CLIVM_EXPERIMENTAL_X_SYS_EXEC")

func (runner *Runner) RunXSysExec(args ...string) error {
	return errXSysNotSuppported
}
