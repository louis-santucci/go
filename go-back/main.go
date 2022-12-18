package main

import (
	"go-go.com/go-back/database"
	_ "go-go.com/go-back/docs"
	"go-go.com/go-back/routes"
)

// SWAGGER CONFIG
// @title           GoApp
// @version         1.0
// @description     A redirection app made in Go with Gin Framework.

// @contact.name   SANTUCCI Louis
// @contact.email  louissantucci1998@gmail.com

// @host      localhost:9090
// @BasePath
func main() {

	// Connect to SQLite DB
	database.ConnectDatabase(db_filename)

	// Creating GIN Router for endpoints
	router := routes.SetupRouter()

	err := router.Run(":9090")
	if err != nil {
		panic(err)
	}
}
