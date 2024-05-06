package router

import (
	"ujikom/pkg/handlers"
	"ujikom/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	r := router.Group("/api")
	AuthRouter(r)
	UserRouter(r)
	UserProfile(r)
	RecipeRouter(r)
	PublicGroup(r)
	return router
}

func AuthRouter(r *gin.RouterGroup) {
	authHandler := handlers.AuthHandler{}
	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
	r.GET("/me", middlewares.Authentication(), authHandler.Me)
}

func UserProfile(r *gin.RouterGroup) {
	profileHandler := handlers.ProfileHandler{}
	profile := r.Group("/profile")
	{
		profile.Use(middlewares.Authentication())
		profile.GET("/", profileHandler.GetMyProfile)
		profile.GET("/all", middlewares.AllowedRole("admin"), profileHandler.GetAllProfiles)
		profile.POST("/", profileHandler.CreateMyProfile)
		profile.PUT("/", profileHandler.UpdateMyProfile)
		// profile.GET("/my-recipes", profileHandler.GetMyRecipes)

		profile.GET("/collection/all", middlewares.AllowedRole("admin"), profileHandler.GetAllCollections)
		profile.GET("/collection", profileHandler.GetMyCollections)
		profile.GET("/collection/:id", profileHandler.GetCollection)
		profile.POST("/collection", profileHandler.CreateCollection)
		profile.PUT("/collection/:id", profileHandler.UpdateCollection)
		profile.DELETE("/collection/:id", profileHandler.DeleteCollection)

		// profile.GET("/stat/total-like", profileHandler.GetTotalLike)
		// 	profile.GET("/stat/total-recipe", profileHandler.GetTotalRecipe)
		// 	profile.GET("/stat/total-review", profileHandler.GetTotalReview)

	}
}

func UserRouter(r *gin.RouterGroup) {
	userHandler := handlers.UserHandler{}
	users := r.Group("/users")
	{
		users.Use(middlewares.Authentication(), middlewares.AllowedRole("admin"))
		users.GET("/", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)

		users.GET("/:id/recipes", userHandler.GetRecipesByUser)
		users.GET("/:id/reviews", userHandler.GetReviewsByUser)
		// users.GET("/:id/likes", userHandler.GetLikesByUser)
	}
}

func RecipeRouter(r *gin.RouterGroup) {
	recipeHandler := handlers.RecipeHandler{}
	IngredientHandler := handlers.IngredientHandler{}
	profileHandler := handlers.ProfileHandler{}
	recipes := r.Group("/recipes")
	{
		recipes.Use(middlewares.Authentication())
		recipes.POST("/", recipeHandler.CreateRecipe)
		recipes.GET("/my", profileHandler.GetMyRecipes)
		recipes.GET("/", middlewares.AllowedRole("admin"), recipeHandler.GetRecipes)
		recipes.GET("/:id", recipeHandler.GetRecipe)
		recipes.PUT("/:id", recipeHandler.UpdateRecipe)
		recipes.DELETE("/:id", recipeHandler.DeleteRecipe)

		recipes.POST("/:id/like", recipeHandler.LikeRecipe)
		recipes.POST("/:id/add-to-collection", recipeHandler.AddToCollection)

		recipes.POST("/:id/ingredients", IngredientHandler.AddIngredient)
		recipes.PUT("/:id/ingredients/:ingredient_id", IngredientHandler.UpdateIngredient)
		recipes.DELETE("/:id/ingredients/:ingredient_id", IngredientHandler.DeleteIngredient)

		recipes.GET("/:id/reviews", recipeHandler.GetReviews)
		recipes.POST("/:id/reviews", recipeHandler.CreateReview)
		recipes.PUT("/:id/reviews/:review_id", recipeHandler.UpdateReview)
		recipes.DELETE("/:id/reviews/:review_id", recipeHandler.DeleteReview)

		recipes.POST("/tags", recipeHandler.CreateTag)
		recipes.GET("/tags", recipeHandler.GetTags)
		recipes.DELETE("/tags/:id", recipeHandler.DeleteTag)

		recipes.POST("/labels", recipeHandler.CreateLabel)
		recipes.GET("/labels", recipeHandler.GetLabels)
		recipes.DELETE("/labels/:id", recipeHandler.DeleteLabel)

		// recipes.POST("/helpfuls", recipeHandler.CreateHelpful)
	}
}

func PublicGroup(r *gin.RouterGroup) {
	recipeHandler := handlers.RecipeHandler{}
	publicGroup := r.Group("/public")
	{
		publicGroup.GET("/recipes", recipeHandler.GetPublicRecipes)
		publicGroup.GET("/recipes/:id", recipeHandler.GetRecipe)
		publicGroup.GET("/recipes/:id/reviews", recipeHandler.GetReviews)
		publicGroup.GET("/search", recipeHandler.SearchRecipe)
		publicGroup.GET("/collections", recipeHandler.GetPublicCollections)
	}
}
