package action_replace

import (
	"strings"
)

type Action struct{}

func (a *Action) Perform(input string, options string) string {
	opt := strings.Split(options, "|")
	output := strings.ReplaceAll(input, opt[0], opt[1])
	return output
}
