package actions

import (
	"github.com/rumikk/actgo/actions/action_get"
	"github.com/rumikk/actgo/actions/action_log"
	"github.com/rumikk/actgo/actions/action_notification"
	"github.com/rumikk/actgo/actions/action_prefix"
	"github.com/rumikk/actgo/actions/action_replace"
	"github.com/rumikk/actgo/actions/action_suffix"
	"github.com/rumikk/actgo/actions/action_transmission"
)

type Action interface {
	Perform(input string, options any) string
}

var actions = map[string]func() Action{
	"get":          func() Action { return &action_get.Action{} },
	"log":          func() Action { return &action_log.Action{} },
	"notification": func() Action { return &action_notification.Action{} },
	"prefix":       func() Action { return &action_prefix.Action{} },
	"replace":      func() Action { return &action_replace.Action{} },
	"suffix":       func() Action { return &action_suffix.Action{} },
	"transmission": func() Action { return &action_transmission.Action{} },
}

func NewAction(action string) Action {
	return actions[action]()
}
