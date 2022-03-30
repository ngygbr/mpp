package main

import (
	"fmt"
	"mpp/internal/api"
	utils "mpp/pkg/config"
	"mpp/pkg/controller"
	"mpp/pkg/db"
)

func main() {

	xd := utils.GetConfig()
	fmt.Println(xd)

	if _, err := db.Connect("/tmp"); err != nil {
		panic("failed to connect database")
	}

	go controller.SetSettledAfter24Hours()

	if err := api.Init(); err != nil {
		panic("failed init")
	}

}
