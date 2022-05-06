package main

import (
	"github.com/ngygbr/mpp/internal/api"
	utils "github.com/ngygbr/mpp/pkg/config"
	"github.com/ngygbr/mpp/pkg/controller"
	"github.com/ngygbr/mpp/pkg/db"
)

func main() {

	utils.LogConfig()
	config := utils.GetConfig()

	if _, err := db.Connect(config.BadgerTmp); err != nil {
		panic("failed to connect database")
	}

	go controller.SetSettledAfter24Hours()

	if err := api.Init(); err != nil {
		panic("failed init")
	}

}
