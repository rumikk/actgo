package action_get

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

type Action struct{}

func (a *Action) Perform(input string, options string) string {
	_, err := url.Parse(options)
	if err != nil {
		log.Fatal("action_get: error when parsing url")
	}
	resp, err := http.Get(options)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal("action_get:", err.Error())
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("action_get:", err.Error())
	}
	return string(b)
}
