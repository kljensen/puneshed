package menu

import (
	"io"
	"io/ioutil"

	"github.com/kljensen/puneshed/shell"
	"gopkg.in/yaml.v2"
)

// Action -- a form of punishment or other
// action. Each action has a human description
// and a shell command that gets executed for
// that.
type Action struct {
	Description string
	Command     string
	Args        []string
}

func (a Action) run() ([]byte, error) {
	return shell.Run(a.Command, a.Args...)
}

// ActionList -- a list of actions
type ActionList struct {
	Description string
	Actions     []Action
}

// Unmarshall from a reader object
func Unmarshall(r io.Reader) (ActionList, error) {
	actions := ActionList{}
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return actions, err
	}
	err = yaml.Unmarshal(buf, &actions)
	return actions, err
}
