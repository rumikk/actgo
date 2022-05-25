package action_get

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

type Action struct {
	Url       string `yaml:"url"`
	UserAgent string `yaml:"useragent"`
}

func (a *Action) Perform(input string, options any) string {
	o := options.(*Action)
	if input != "" && o.Url == "" {
		o.Url = input
	}
	parsed, err := url.Parse(o.Url)
	if err != nil {
		log.Fatal("action_get: error when parsing url")
	}
	resp, err := get(parsed.String(), o.UserAgent)
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

func get(url string, useragent string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", useragent)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
