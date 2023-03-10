package controllers

import (
	"github.com/gin-gonic/gin"
	"louissantucci/goapi/database"
	"louissantucci/goapi/models"
	"louissantucci/goapi/responses"
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
	var historyResults []models.HistoryEntry
	database.DB.Find(&historyResults)
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
	err := database.DB.Exec("DELETE FROM history_entries").Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	responseData := "History purged"
	c.JSON(http.StatusOK, responses.NewOKResponse(responseData))
}
