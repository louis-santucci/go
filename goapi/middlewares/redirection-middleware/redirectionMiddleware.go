package redirection_middleware

import (
	"github.com/gin-gonic/gin"
	"louissantucci/goapi/database"
	"louissantucci/goapi/models"
	"net/http"
	"strings"
)

// HandleCall
// How it works: This middleware handles all requests made for the API
// - First, if the call begins with '/api', it will let the request pass
// - Else, the call is for a redirection (go /xxx):
// - check if xxx is a valid alias
// - if yes: redirect
// - else: redirect to front-end to create redirection
func HandleCall(c *gin.Context) {
	url := c.Request.URL.Path
	// if url begins with /api, next
	if strings.HasPrefix(url, "/api") || strings.Compare(url, "/favicon.ico") == 0 {
		c.Next()
	} else {
		HandleRedirectionCall(c)
	}
}

func HandleRedirectionCall(c *gin.Context) {
	var redirection models.Redirection
	shortcut := c.Request.URL.Path
	err := database.DB.Where("shortcut = ?", shortcut).First(&redirection).Error
	if err != nil {
		c.Redirect(http.StatusSeeOther, "http://localhost:9091/#/error/notFound")
		return
	}
	newUrl := redirection.RedirectUrl

	// Incrementation
	redirection.Views = redirection.Views + 1

	database.DB.Model(&redirection).Updates(redirection)

	c.Redirect(http.StatusSeeOther, newUrl)

	return
}
