package router

import (
	"ujikom/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	r := router.Group("/api")
	AuthRouter(r)
	UserRouter(r)
	return router
}

func AuthRouter(r *gin.RouterGroup) {
	authHandler := handlers.AuthHandler{}
	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
}

func UserRouter(r *gin.RouterGroup) {
	userHandler := handlers.UserHandler{}
	users := r.Group("/users")
	{
		users.GET("/", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}
}
