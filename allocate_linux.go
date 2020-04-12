// +build linux

package chromedp

import (
	"os/exec"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

func allocateCmdOptions(cmd *exec.Cmd) {
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = new(syscall.SysProcAttr)
	}

	if !isJailed() {
		// When the parent process dies (Go), kill the child as well.
		cmd.SysProcAttr.Pdeathsig = syscall.SIGKILL
	}
}

func isJailed() bool {
	var sig int

	err := unix.Prctl(unix.PR_GET_PDEATHSIG, uintptr(unsafe.Pointer(&sig)), 0, 0, 0)
	if err != nil {
		return err.Error() == "operation not permitted"
	}

	return false

}
