package action_suffix

type Action struct {
	Text string `yaml:"text"`
}

func (a *Action) Perform(input string, options any) string {
	o := options.(*Action)
	return input + o.Text
}
