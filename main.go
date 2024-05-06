package main

import (
	"ujikom/database"
	"ujikom/pkg/router"
	"ujikom/config"
)

func Init(){
	config, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	database.Init(&config)
}

func main() {
	Init()
	r := router.SetupRouter()
	r.Run(":8080")
}
