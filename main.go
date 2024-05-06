package main

import (
	"ujikom/config"
	"ujikom/database"
	"ujikom/pkg/router"
)

func Init() {
	config := config.LoadConfig()
	database.Init(&config)
}

func main() {
	Init()
	r := router.SetupRouter()
	r.Run(":8080")
}
