package config

import (
	"github.com/gin-contrib/cors"
)

func CorsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:9091", "https://go.go", "http://go.go"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Origin"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AddAllowMethods("OPTIONS", "PUT", "POST", "GET", "DELETE")
	return corsConfig
}
