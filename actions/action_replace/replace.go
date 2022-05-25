package action_replace

import (
	"strings"
)

type Action struct {
	Old string `yaml:"old"`
	New string `yaml:"new"`
}

func (a *Action) Perform(input string, options any) string {
	o := options.(*Action)
	output := strings.ReplaceAll(input, o.Old, o.New)
	return output
}
