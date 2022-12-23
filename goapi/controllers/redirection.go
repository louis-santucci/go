package controllers

import (
	"github.com/gin-gonic/gin"
	"louissantucci/goapi/database"
	"louissantucci/goapi/models"
	"louissantucci/goapi/responses"
	"net/http"
	"time"
)

// GET /redirection

// GetRedirections 				godoc
// @Summary 					Get all redirections
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Success						200 	{object} 	responses.OKResponse
// @Router						/redirection [get]
func GetRedirections(c *gin.Context) {
	var redirections []models.Redirection

	database.DB.Find(&redirections)

	c.JSON(http.StatusOK, responses.NewOKResponse(redirections))
}

// GET /redirection/:id

// GetRedirection 				godoc
// @Summary 					Get one redirection in function of its ID
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Param						id 		path		int true "id"
// @Success						200 	{object} 	responses.OKResponse
// @Failure						404 	{object} 	responses.ErrorResponse
// @Router						/redirection/{id} [get]
func GetRedirection(c *gin.Context) {
	var redirection models.Redirection

	err := database.DB.Where("id = ?", c.Param("id")).First(&redirection).Error
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewOKResponse(redirection))
}

// PUT /redirection/:id

// IncrementRedirectionView 	godoc
// @Summary 					Increments number of view for one redirection in fucntion of its ID
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Param						id 		path		int true "id"
// @Success						200 	{object} 	responses.OKResponse
// @Failure						404 	{object} 	responses.ErrorResponse
// @Router						/redirection/{id} [put]
func IncrementRedirectionView(c *gin.Context) {
	var redirection models.Redirection

	err := database.DB.Where("id = ?", c.Param("id")).First(&redirection).Error
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	// Incrementation
	redirection.Views = redirection.Views + 1

	database.DB.Model(&redirection).Updates(redirection)

	c.JSON(http.StatusOK, responses.NewOKResponse(redirection))
}

// POST /redirection/:id

// EditRedirection 			godoc
// @Summary						Updates a redirection in function of its ID
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Param						id 		path		int true "id"
// @Param						request body models.RedirectionInput true "query params"
// @Success						200 	{object} 	responses.OKResponse
// @Failure						400 	{object} 	responses.ErrorResponse
// @Failure						404 	{object} 	responses.ErrorResponse
// @Failure						500 	{object} 	responses.ErrorResponse
// @Router						/redirection/{id} [post]
func EditRedirection(c *gin.Context) {
	var redirection models.Redirection

	err := database.DB.Where("id = ?", c.Param("id")).First(&redirection).Error
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	// Input validation
	var input models.RedirectionInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}
	redirection.Shortcut = input.Shortcut
	redirection.RedirectUrl = input.RedirectUrl
	redirection.UpdatedAt = time.Now()
	err = database.DB.Model(&redirection).Updates(redirection).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewOKResponse(redirection))
}

// POST /redirection

// CreateRedirection			godoc
// @Summary 					Creates a new redirection
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Param						request body models.RedirectionInput true "query params"
// @Success						200 	{object} 	responses.OKResponse
// @Failure						400 	{object} 	responses.ErrorResponse
// @Failure						500 	{object} 	responses.ErrorResponse
// @Router						/redirection [post]
func CreateRedirection(c *gin.Context) {
	// Input validation
	var input models.RedirectionInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// Redirection creation
	redirection := models.Redirection{
		Shortcut:    input.Shortcut,
		RedirectUrl: input.RedirectUrl,
		Views:       0,
		CreatedAt:   time.Now(),
	}
	err = database.DB.Create(&redirection).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.NewOKResponse(redirection))
}

// DELETE /redirection/:id

// DeleteRedirection			godoc
// @Summary 					Deletes a redirection in function of its ID
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Param						id 		path		int true "id"
// @Success						200 	{object} 	responses.OKResponse
// @Failure						404 	{object} 	responses.ErrorResponse
// @Failure						500 	{object} 	responses.ErrorResponse
// @Router						/redirection/{id} [delete]
func DeleteRedirection(c *gin.Context) {
	// Gets model if exists
	var redirection models.Redirection
	id := c.Param("id")
	err := database.DB.Where("id = ?", id).First(&redirection).Error
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	err = database.DB.Delete(&redirection).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	responseData := "Redirection #" + id + " deleted"
	c.JSON(http.StatusOK, responses.NewOKResponse(responseData))
}
