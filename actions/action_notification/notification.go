package action_notification

import "github.com/0xAX/notificator"

type Action struct{}

func (a *Action) Perform(input string, options string) string {
	notify := notificator.New(notificator.Options{
		DefaultIcon: "",
		AppName:     "ActGo",
	})
	notify.Push("ActGo", input, "", "normal")
	return input
}
