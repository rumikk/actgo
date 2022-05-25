package action_notification

import "github.com/0xAX/notificator"

type Action struct {
	Text string `yaml:"text"`
}

func (a *Action) Perform(input string, options any) string {
	o := options.(*Action)
	if input != "" && o.Text == "" {
		o.Text = input
	}
	notify := notificator.New(notificator.Options{
		DefaultIcon: "",
		AppName:     "Act",
	})
	notify.Push("Act", o.Text, "", "normal")
	return o.Text
}
