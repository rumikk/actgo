package main

import (
	"actgo/parser"
	"actgo/storage"
	"fmt"
	bolt "go.etcd.io/bbolt"
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
}
