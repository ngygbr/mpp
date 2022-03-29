package db

import (
	"log"

	"github.com/dgraph-io/badger/v3"
)

var Database *badger.DB

func Connect() (*badger.DB, error) {
	options := badger.DefaultOptions("/tmp")

	var err error
	Database, err = badger.Open(options)
	if err != nil {
		log.Fatal(err)
	}

	return Database, nil
}
