package actions

import (
	"actgo/actions/action_log"
	"actgo/actions/action_notification"
	"actgo/actions/action_replace"
	"actgo/actions/action_transmission"
)

type Action interface {
	Perform(input string, options string) string
}

var actions = map[string]func() Action{
	"log":          func() Action { return &action_log.Action{} },
	"replace":      func() Action { return &action_replace.Action{} },
	"transmission": func() Action { return &action_transmission.Action{} },
	"notification": func() Action { return &action_notification.Action{} },
}

func NewAction(action string) Action {
	return actions[action]()
}
