package app

import (
	"bytes"
	"github.com/wonderivan/logger"
	"os"
	"os/exec"
	"strings"
)

func Cmd(name string, arg ...string) {
	cmd := exec.Command(name, arg[:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		logger.Error("调用执行失败。", err)
	}
}

func CmdString(name string, arg ...string) string {
	cmd := exec.Command(name, arg[:]...)
	cmd.Stdin = os.Stdin
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	err := cmd.Run()
	if err != nil {
		logger.Error("调用执行失败。", err)
		return ""
	}
	return b.String()
}

func CmdVagrantMachine() []string {
	vms := CmdString("/bin/sh", "-c", "vagrant status  | grep virtualbox | awk '{print $1}'")
	return strings.Split(vms, "\n")
}
