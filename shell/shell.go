package shell

import (
	"os/exec"
)

var execCommand = exec.Command

// Run - echo a command
func Run(s string) ([]byte, error) {
	cmd := execCommand("echo", s)
	out, err := cmd.CombinedOutput()
	return out, err
}
