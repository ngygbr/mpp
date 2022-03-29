package db

import (
	"log"

	"github.com/dgraph-io/badger/v3"
)

var Database *badger.DB

func Connect(path string) (*badger.DB, error) {
	options := badger.DefaultOptions(path)

	var err error
	Database, err = badger.Open(options)
	if err != nil {
		log.Fatal(err)
	}

	return Database, nil
}
