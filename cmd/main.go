package main

import (
	"actgo/parser"
	"fmt"
	"log"
	"os"
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

	processParser, err := parser.NewProcessParser(processFile)
	if err != nil {
		log.Fatal(err)
	}
}
