package shell

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

const echoResult = "foo!"

func TestRun(t *testing.T) {
	// Notice that `execCommand` is a package-level global
	// that is being overridden here.
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()
	out, err := Run("wooooot")
	if err != nil {
		t.Errorf("Expected nil error, got %#v", err)
	}
	if string(out) != echoResult {
		t.Errorf("Expected %q, got %q", echoResult, out)
	}
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	// some code here to check arguments perhaps?
	fmt.Fprintf(os.Stdout, echoResult)
	os.Exit(0)
}
