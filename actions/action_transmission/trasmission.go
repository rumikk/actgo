package action_transmission

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Action struct{}

func (a *Action) Perform(input string, options string) string {
	client := &http.Client{}

	getSessionForm := url.Values{"method": {"session-get"}}

	addForm := url.Values{
		"method":    {"torrent-add"},
		"arguments": {"filename", input},
	}

	req, err := http.NewRequest("POST", "http://localhost:9091/transmission/rpc", strings.NewReader(getSessionForm.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	sessionId := resp.Header.Get("X-Transmission-Session-Id")

	req, err = http.NewRequest("POST", "http://localhost:9091/transmission/rpc", strings.NewReader(addForm.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Transmission-Session-Id", sessionId)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	return input
}
