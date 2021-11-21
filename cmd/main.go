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
			selected := selection.Text()
			if selected != "" {
				extracted := selection.AttrOr("href", "")
				if extracted != "" {
					entry := &storage.Entry{
						Name:      process.Name,
						Selected:  selected,
						Extracted: extracted,
					}
					if s.FindEntry(entry.Extracted).Extracted == "" {
						s.AddEntry(entry)
						for _, action := range process.Actions {
							newAction := actions.NewAction(action.Name)
							extracted = newAction.Perform(extracted, action.Options)
						}
					}
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
