package main

import (
	"mpp/internal/api"
	utils "mpp/pkg/config"
	"mpp/pkg/controller"
	"mpp/pkg/db"
)



func main() {

	utils.LogConfig()

	if _, err := db.Connect("/tmp"); err != nil {
		panic("failed to connect database")
	}

	go controller.SetSettledAfter24Hours()

	if err := api.Init(); err != nil {
		panic("failed init")
	}

}
