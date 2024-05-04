package main

import "github.com/gin-gonic/gin"

func main() {
	gin := gin.Default()

	gin.Run(":8080")
}
