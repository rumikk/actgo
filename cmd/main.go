package main

import (
	"actgo/actions"
	"actgo/parser"
	"actgo/storage"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	bolt "go.etcd.io/bbolt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: act [db] [config]")
		os.Exit(1)
	}

	processFile, err := os.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	processes, err := parser.NewProcessParser(processFile)
	if err != nil {
		log.Fatal(err)
	}

	db, err := bolt.Open(os.Args[1], 0600, &bolt.Options{})
	if err != nil {
		log.Fatal(err)
	}

	s := storage.Storage{Db: db}
	s.Init()

	for _, process := range processes {
		page, err := GetPage(process.Url)
		if err != nil {
			log.Println(err)
			continue
		}
		page.Find(strings.TrimSpace(process.Selector)).Each(func(i int, selection *goquery.Selection) {
			input := selection.AttrOr("href", "")
			if input != "" {
				if s.FindEntry(input).Input == "" {
					var output string
					output = input
					for _, action := range process.Actions {
						newAction := actions.NewAction(action.Name)
						output = newAction.Perform(output, action.Options)
					}
					s.AddEntry(&storage.Entry{
						Name:   process.Name,
						Input:  input,
						Output: output,
					})
				}
			}
		})
	}
}

func GetPage(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
