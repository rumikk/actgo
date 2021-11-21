package action_print

import "fmt"

type Action struct{}

func (a *Action) Perform(extracted string) string {
	fmt.Println(extracted)
	return extracted
}
