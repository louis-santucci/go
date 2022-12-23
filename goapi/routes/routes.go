package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"louissantucci/goapi/config"
	"louissantucci/goapi/controllers"
	"louissantucci/goapi/middlewares/jwt"
	"louissantucci/goapi/responses"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Swagger

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Redirections

	api := router.Group("/api")
	{
		redirection := api.Group("/redirection")
		{
			redirection.GET("", controllers.GetRedirections)
			redirection.GET("/:id", controllers.GetRedirection)
			redirection.POST("", jwt.JWTTokenCheck, controllers.CreateRedirection)
			redirection.POST("/:id", jwt.JWTTokenCheck, controllers.EditRedirection)
			redirection.PUT("/:id", controllers.IncrementRedirectionView)
			redirection.PATCH("/:id", jwt.JWTTokenCheck, controllers.ResetRedirectionView)
			redirection.DELETE("/:id", jwt.JWTTokenCheck, controllers.DeleteRedirection)
		}

		user := api.Group("/user")
		{
			user.POST("/login", controllers.LoginUser)
			user.POST("/register", controllers.RegisterUser)
			user.POST("/edit/:id", jwt.JWTTokenCheck, controllers.EditUser)
		}
	}

	// No routes

	router.NoRoute(func(c *gin.Context) {
		errorData := "Route NOT FOUND"
		c.JSON(http.StatusNotFound, responses.NewErrorResponse(http.StatusNotFound, errorData))
	})

	// CORS Config
	router.Use(cors.New(config.CorsConfig()))

	return router
}
