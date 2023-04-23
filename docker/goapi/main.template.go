package main

import (
	"louissantucci/goapi/config"
	"louissantucci/goapi/database"
	_ "louissantucci/goapi/docs"
	"louissantucci/goapi/routes"
)

// SWAGGER CONFIG
// @title           GoApp
// @version         1.0
// @description     A redirection app made in Go with Gin Framework.

// @contact.name   	SANTUCCI Louis
// @contact.email  	louissantucci1998@gmail.com

// @host      		${IP}:9090
// @BasePath		/api
func main() {

	// Connect to SQLite DB
	database.ConnectDatabase(config.Db_filename)

	// Creating GIN Router for endpoints
	router := routes.SetupRouter()

	var err error
	if config.TLS_enabled {
		err = router.RunTLS(":9090", "./certs/nginx.crt", "./certs/nginx.key")
	} else {
		err = router.Run(":9090")
	}
	if err != nil {
		panic(err)
	}
}
