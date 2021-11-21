package action_print

import (
	"fmt"
)

type Action struct{}

func (a *Action) Perform(input string, options string) string {
	output := options + input
	fmt.Println(output)
	return output
}
