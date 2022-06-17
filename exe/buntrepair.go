package main

import (
	"os"
	"log"
	"github.com/tidwall/buntdb"
)

func main() {
	file := os.Args[1]
	db, err := buntdb.Open(file)
	if err != nil {
		log.Fatalf("failed to load: %v", err)
	}

	err = db.Shrink()
	if err != nil {
		log.Fatalf("failed to rebuild: %v", err)
	}

	err = db.Close()
	if err != nil {
		log.Fatalf("failed to close: %v", err)
	}
}
