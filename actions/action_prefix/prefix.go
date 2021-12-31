package action_prefix

type Action struct{}

func (a *Action) Perform(input string, options string) string {
	return options + input
}
