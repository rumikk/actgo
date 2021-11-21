package action_replace

import (
	"log"
	"strings"
)

type Action struct{}

func (a *Action) Perform(input string, options string) string {
	opt := strings.Split(options, "|")
	if len(opt) < 2 {
		log.Fatal("Error when processing input: " + input + " options: " + options)
	}
	output := strings.ReplaceAll(input, opt[0], opt[1])
	return output
}
