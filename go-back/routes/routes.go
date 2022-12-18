package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-go.com/go-back/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Swagger

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Redirections

	api := router.Group("/api")
	{
		api.GET("/redirection", controllers.GetRedirections)
		api.GET("/redirection/:id", controllers.GetRedirection)
		api.POST("/redirection", controllers.CreateRedirection)
		api.POST("/redirection/:id", controllers.UpdateRedirection)
		api.PUT("/redirection/:id", controllers.IncrementRedirectionView)
	}

	// No routes

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "NOT_FOUND", "message": "Not found"})
	})

	return router
}
