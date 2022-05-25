package action_log

import (
	"log"
)

type Action struct{}

func (a *Action) Perform(input string, options any) string {
	log.Println(input)
	return input
}
