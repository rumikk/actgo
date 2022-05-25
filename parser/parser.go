package parser

import (
	"errors"
	"github.com/rumikk/actgo/actions/action_get"
	"github.com/rumikk/actgo/actions/action_log"
	"github.com/rumikk/actgo/actions/action_notification"
	"github.com/rumikk/actgo/actions/action_prefix"
	"github.com/rumikk/actgo/actions/action_replace"
	"github.com/rumikk/actgo/actions/action_suffix"
	"github.com/rumikk/actgo/actions/action_transmission"
	"gopkg.in/yaml.v3"
)

type Actions struct {
	Name    string `yaml:"name"`
	Actions []Action
}

type Action struct {
	Name    string      `yaml:"name"`
	Options interface{} `yaml:"-"`
}

func (action *Action) UnmarshalYAML(n *yaml.Node) error {
	type A Action
	type T struct {
		*A      `yaml:",inline"`
		Options yaml.Node `yaml:"options"`
	}
	obj := &T{A: (*A)(action)}
	err := n.Decode(obj)
	if err != nil {
		return err
	}

	switch action.Name {
	case "get":
		action.Options = new(action_get.Action)
	case "log":
		action.Options = new(action_log.Action)
	case "notification":
		action.Options = new(action_notification.Action)
	case "prefix":
		action.Options = new(action_prefix.Action)
	case "replace":
		action.Options = new(action_replace.Action)
	case "suffix":
		action.Options = new(action_suffix.Action)
	case "transmission":
		action.Options = new(action_transmission.Action)
	default:
		return errors.New("Unknown action!")
	}

	return obj.Options.Decode(action.Options)
}

func Parse(file []byte) ([]Actions, error) {
	var actions []Actions
	err := yaml.Unmarshal(file, &actions)
	if err != nil {
		return nil, err
	}
	return actions, nil
}
