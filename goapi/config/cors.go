package config

import (
	"github.com/gin-contrib/cors"
)

const AUTHORIZATION = "Authorization"
const ORIGIN = "Origin"
const CONTENT_TYPE = "Content-Type"
const CONTENT_LENGTH = "Content-Length"

func CorsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	const host = HOST + ":9091"
	corsConfig.AllowOrigins = []string{"http://" + host, "https://" + host}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{ORIGIN, AUTHORIZATION, CONTENT_TYPE}
	corsConfig.ExposeHeaders = []string{CONTENT_TYPE, CONTENT_LENGTH}
	corsConfig.AddAllowMethods("OPTIONS", "PUT", "POST", "GET", "DELETE")
	return corsConfig
}
