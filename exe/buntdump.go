package main

import (
	"os"
	"fmt"
	"log"
	"github.com/tidwall/buntdb"
)

func main() {
	file := os.Args[1]
	db, err := buntdb.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			ttl, err := tx.TTL(key)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("key: %s, value: %s, ttl: %v\n", key, value, ttl)
			return true
		})
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}
