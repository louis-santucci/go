package config

import (
	"github.com/gin-contrib/cors"
)

const AUTHORIZATION = "Authorization"
const ORIGIN = "Origin"
const CONTENT_TYPE = "Content-Type"
const ANGOLAR_SECRET = "ANGOLAR_SECRET"
const CONTENT_LENGTH = "Content-Length"

func CorsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:9091", "https://localhost:9091", "https://go.go", "http://go.go"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{ORIGIN, AUTHORIZATION, CONTENT_TYPE, ANGOLAR_SECRET}
	corsConfig.ExposeHeaders = []string{CONTENT_TYPE, CONTENT_LENGTH}
	corsConfig.AddAllowMethods("OPTIONS", "PUT", "POST", "GET", "DELETE")
	return corsConfig
}
