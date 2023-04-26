package controllers

import (
	"github.com/gin-gonic/gin"
	"louissantucci/goapi/responses"
	"louissantucci/goapi/services"
	"net/http"
)

// GetHistory			godoc
// @Summary 					Gets history of redirection executions
// @Tags						history
// @Accept						json
// @Produce						json
// @Success						200 	    {object} 	responses.OKResponse
// @Router						/history    [get]
func GetHistory(c *gin.Context) {
	historyResults, err := services.GetHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	c.JSON(http.StatusOK, responses.NewOKResponse(historyResults))
}

// ResetHistory			godoc
// @Summary 					Gets history of redirection executions
// @Tags						history
// @Accept						json
// @Produce						json
// @Success						200 	    {object} 	responses.OKResponse
// @Failure						500         {object} 	responses.ErrorResponse
// @Router						/history/reset    [delete]
func ResetHistory(c *gin.Context) {
	err := services.ResetHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	responseData := "History purged"
	c.JSON(http.StatusOK, responses.NewOKResponse(responseData))
}
