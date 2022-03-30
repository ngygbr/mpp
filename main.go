package main

import (
	"mpp/internal/api"
	"mpp/pkg/controller"
	"mpp/pkg/db"
)

func main() {

	if _, err := db.Connect("/tmp"); err != nil {
		panic("failed to connect database")
	}

	go controller.SetSettledAfter24Hours()

	if err := api.Init(); err != nil {
		panic("failed init")
	}

}
