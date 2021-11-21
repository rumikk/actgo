package actions

import "actgo/actions/action_print"

type Action interface {
	Perform(input string) string
}

var actions = map[string]func() Action{
	"print": func() Action { return &action_print.Action{} },
}

func NewAction(action string) Action {
	return actions[action]()
}
