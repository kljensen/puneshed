package menu
import (
	"gopkg.in/yaml.v2"
)

// Action -- a form of punishment or other
// action. Each action has a human description
// and a shell command that gets executed for
// that.
type Action struct {
	Description string
	Command     string
}

func Load
