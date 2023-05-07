package main

import (
	"os"
	"os/exec"
	"strings"
)

func ExecNoOut(args ...string) error {
	return cmdExecNoOut("tmux", args...)
}

func Exec(args ...string) ([]byte, error) {
	return cmdExec("tmux", args...)
}

func cmdExec(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = strings.NewReader("")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, err
	}
	return output, nil
}

func cmdExecNoOut(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	defer func() {
		if cmd.Process != nil {
			cmd.Process.Release()
		}
	}()
	return cmd.Run()
}
