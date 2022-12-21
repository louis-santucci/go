package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"louissantucci/goapi/config"
	"louissantucci/goapi/controllers"
	"louissantucci/goapi/errors"
	"louissantucci/goapi/middlewares/jwt"
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
			redirection.POST("", controllers.CreateRedirection, jwt.JWTTokenCheck)
			redirection.POST("/:id", controllers.UpdateRedirection, jwt.JWTTokenCheck)
			redirection.PUT("/:id", controllers.IncrementRedirectionView)
			redirection.DELETE("/:id", controllers.DeleteRedirection, jwt.JWTTokenCheck)
		}

		user := api.Group("/user")
		{
			user.POST("/login", controllers.LoginUser)
			user.POST("/register", controllers.RegisterUser)
			user.POST("/edit/:id", controllers.EditUser, jwt.JWTTokenCheck)
		}
	}

	// No routes

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": errors.NotFoundError, "message": "Not found"})
	})

	// CORS Config
	router.Use(cors.New(config.CorsConfig()))

	return router
}
