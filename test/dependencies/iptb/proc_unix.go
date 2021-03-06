// +build !windows

package main

import (
	"os/exec"
	"syscall"
)

func init() {
	setupOpt = func(cmd *exec.Cmd) {
		cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	}
}
