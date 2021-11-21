package actions

import (
	"actgo/actions/action_print"
	"actgo/actions/action_replace"
)

type Action interface {
	Perform(input string, options string) string
}

var actions = map[string]func() Action{
	"print":   func() Action { return &action_print.Action{} },
	"replace": func() Action { return &action_replace.Action{} },
}

func NewAction(action string) Action {
	return actions[action]()
}
