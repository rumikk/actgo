package action_prefix

type Action struct {
	Text string `yaml:"text"`
}

func (a *Action) Perform(input string, options any) string {
	o := options.(*Action)
	return o.Text + input
}
