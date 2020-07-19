package menu

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var stringMenu = `
description: My fake menu 
actions:
    - description: echo foo
      command: echo
      args:
          - foo
    - description: echo bar
      command: echo
      args:
          - bar
`
var expectedMenu = ActionList{
	Description: "My fake menu",
	Actions: []Action{
		{
			Description: "echo foo",
			Command:     "echo",
			Args: []string{
				"foo",
			},
		},
		{
			Description: "echo bar",
			Command:     "echo",
			Args: []string{
				"bar",
			},
		},
	},
}

func TestUnmarshall(t *testing.T) {
	unmarshalledMenu, err := Unmarshall(strings.NewReader(stringMenu))
	if err != nil {
		t.Errorf("Expected nil error, got %#v", err)
	}
	if cmp.Equal(unmarshalledMenu, expectedMenu) == false {
		t.Error("Incorrectly unmarshalled menu")
	}
}
