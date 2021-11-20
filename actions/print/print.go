package print

import "fmt"

type Action struct{}

func (a *Action) Perform(extracted string) {
	fmt.Println(extracted)
}
