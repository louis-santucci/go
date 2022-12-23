package middlewares

import (
	"github.com/gin-gonic/gin"
	"louissantucci/goapi/config"
	"louissantucci/goapi/responses"
	"net/http"
)

func AngolarTokenCheck(c *gin.Context) {
	angolarToken := c.GetHeader("AngolarToken")
	if angolarToken == "" {
		errorData := "AngolarToken is empty"
		c.AbortWithStatusJSON(http.StatusForbidden, responses.NewErrorResponse(http.StatusForbidden, errorData))
		return
	}

	if angolarToken != config.Angolar_secret {
		errorData := "AngolarToken is invalid"
		c.AbortWithStatusJSON(http.StatusForbidden, responses.NewErrorResponse(http.StatusForbidden, errorData))
		return
	}

	c.Next()
}
