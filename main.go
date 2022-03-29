package main

import (
	"mock-paymentprocessor/internal/api"
	"mock-paymentprocessor/pkg/db"
	"mock-paymentprocessor/pkg/transaction"
)

func main() {

	if _, err := db.Connect(); err != nil {
		panic("failed to connect database")
	}

	go transaction.SetSettledAfter24Hours()

	if err := api.Init(); err != nil {
		panic("failed init")
	}

}
