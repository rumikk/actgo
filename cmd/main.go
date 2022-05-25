package main

import (
	"fmt"
	"github.com/rumikk/actgo/actions"
	"github.com/rumikk/actgo/parser"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: act [Actfile] [DB]")
		os.Exit(1)
	}

	actFile, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	jobs, err := parser.Parse(actFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, job := range jobs {
		var next string
		for _, action := range job.Actions {
			a := actions.NewAction(action.Name)
			next = a.Perform(next, action.Options)
		}
	}
}
