package handlers

import (
	"ujikom/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type AppHandler struct {
}

func NewAppHandler() AppHandler {
	return AppHandler{}
}

func (h *AppHandler) Run(c *gin.Context) {
	appData := map[string]interface{}{
		"app":        "ujikom",
		"version":    "1.0.0",
		"author":     "Indra Faizal Amri",
		"go_version": "1.21.6",
		"framework":  "Gin Gonic",
		"databases":  "PostgreSQL",
		"message":    "Welcome to Ujikom API!",
		"context":    "This is a simple API for Ujikom project.",
		"contributors": []string{
			"Indra Faizal Amri",
			"Github Copilot",
			"Google",
			"ChatGPT",
			"Keyboard DST",
			"Laptop HP 14S",
		},
		"note": "Banyak kekurangan bukan karena malas(lumayan banyak sih sebenernya), tapi karena mepet",
	}

	helpers.ResOK(c, appData)
}
