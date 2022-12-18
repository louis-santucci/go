package main

import (
	"github.com/gin-gonic/gin"
	"go.com/models"
	"net/http"
	"time"
)

var redirections = []models.Redirection{
	{ID: 1, SHORTCUT: "fb", REDIRECT_URL: "https://facebook.com", CREATED_AT: time.Now(), VIEWS: 0},
	{ID: 2, SHORTCUT: "fb", REDIRECT_URL: "https://facebook.com", CREATED_AT: time.Now(), VIEWS: 5},
	{ID: 3, SHORTCUT: "fb", REDIRECT_URL: "https://facebook.com", CREATED_AT: time.Now(), VIEWS: 3},
	{ID: 4, SHORTCUT: "fb", REDIRECT_URL: "https://facebook.com", CREATED_AT: time.Now(), VIEWS: 2},
	{ID: 5, SHORTCUT: "fb", REDIRECT_URL: "https://facebook.com", CREATED_AT: time.Now(), VIEWS: 20},
}

func getRedirections(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, redirections)
}

func main() {
	router := gin.Default()
	router.GET("/redirections", getRedirections)
	router.Run("localhost:9090")
}
