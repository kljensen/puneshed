package shell

import (
	"os/exec"
)

var execCommand = exec.Command

// Run - echo a command
func Run(command string, args ...string) ([]byte, error) {
	cmd := execCommand(command, args...)
	out, err := cmd.CombinedOutput()
	return out, err
}
