package action_log

import (
	"log"
)

type Action struct{}

func (a *Action) Perform(input string, options string) string {
	output := options + input
	log.Println(output)
	return output
}
