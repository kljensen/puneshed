package pf

import (
	"os/exec"
)

var execCommand = exec.Command

// Echo - echo a command
func Echo(s string) ([]byte, error) {
    cmd := execCommand("echo", s)
	out, err := cmd.CombinedOutput()
	return out, err
}