package main

import (
	"mpp/internal/api"
	"mpp/pkg/db"
	"mpp/pkg/transaction"
)

func main() {

	if _, err := db.Connect("/tmp"); err != nil {
		panic("failed to connect database")
	}

	go transaction.SetSettledAfter24Hours()

	if err := api.Init(); err != nil {
		panic("failed init")
	}

}
