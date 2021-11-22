package action_transmission

import (
	"log"
	"net/http"
	"strings"
)

type Action struct{}

func (a *Action) Perform(input string, options string) string {
	client := &http.Client{}
	request, err := http.NewRequest("POST", options, strings.NewReader("{\"method\": {\"session-get\"}"))
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	sessionId := response.Header.Get("X-Transmission-Session-Id")
	addTorrentRequest := "{\"method\":\"torrent-add\",\"arguments\":{\"filename\":\"" + input + "\"}}"

	request, err = http.NewRequest("POST", options, strings.NewReader(addTorrentRequest))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("X-Transmission-Session-Id", sessionId)

	response, err = client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return input
}
