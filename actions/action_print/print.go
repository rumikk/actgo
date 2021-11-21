package action_print

import (
	"fmt"
)

type Action struct{}

func (a *Action) Perform(input string) string {
	fmt.Println(input)
	return input
}
