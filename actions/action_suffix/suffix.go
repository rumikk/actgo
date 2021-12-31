package action_suffix

type Action struct{}

func (a *Action) Perform(input string, options string) string {
	return input + options
}
