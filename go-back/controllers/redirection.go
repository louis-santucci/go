package controllers

import (
	"github.com/gin-gonic/gin"
	"go-go.com/go-back/database"
	"go-go.com/go-back/errors"
	"go-go.com/go-back/models"
	"net/http"
	"time"
)

// GET /redirection

// GetRedirections 				godoc
// @Summary 					Get all redirections
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Success						200 	{array} 	models.Redirection
// @Router						/redirection [get]
func GetRedirections(c *gin.Context) {
	var redirections []models.Redirection

	database.DB.Find(&redirections)

	c.JSON(http.StatusOK, gin.H{"data": redirections})
}

// GET /redirection/:id

// GetRedirection 				godoc
// @Summary 					Get one redirection in function of its ID
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Param						id 		path		int true "id"
// @Success						200 	{object} 	models.Redirection
// @Failure						404 	{object} 	error
// @Router						/redirection/{id} [get]
func GetRedirection(c *gin.Context) {
	var redirection models.Redirection

	err := database.DB.Where("id = ?", c.Param("id")).First(&redirection).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.NotFoundError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": redirection})
}

// PUT /redirection/:id

// IncrementRedirectionView 	godoc
// @Summary 					Increments number of view for one redirection in fucntion of its ID
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Param						id 		path		int true "id"
// @Success						200 	{array} 	models.Redirection
// @Failure						404 	{object} 	error
// @Router						/redirection/{id} [put]
func IncrementRedirectionView(c *gin.Context) {
	var redirection models.Redirection

	err := database.DB.Where("id = ?", c.Param("id")).First(&redirection).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.NotFoundError})
		return
	}

	// Incrementation
	var updatedRedirection = models.RedirectionIncrement{Views: redirection.Views + 1}

	database.DB.Model(&redirection).Updates(updatedRedirection)

	c.JSON(http.StatusOK, gin.H{"data": updatedRedirection})
}

// POST /redirection/:id

// UpdateRedirection 			godoc
// @Summary						Updates a redirection in function of its ID
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Param						id 		path		int true "id"
// @Param						request body models.RedirectionInput true "query params"
// @Success						200 	{array} 	models.Redirection
// @Failure						400 	{object} 	error
// @Failure						404 	{object} 	error
// @Router						/redirection/{id} [post]
func UpdateRedirection(c *gin.Context) {
	var redirection models.Redirection

	err := database.DB.Where("id = ?", c.Param("id")).First(&redirection).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.NotFoundError})
		return
	}

	// Input validation
	var input models.RedirectionInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedRedirection := models.Redirection{
		ID:          redirection.ID,
		Shortcut:    input.Shortcut,
		RedirectUrl: input.RedirectUrl,
		Views:       redirection.Views,
		CreatedAt:   redirection.CreatedAt,
		UpdatedAt:   time.Now(),
	}
	database.DB.Model(&redirection).Updates(updatedRedirection)

	c.JSON(http.StatusOK, gin.H{"data": redirection})
}

// POST /redirection

// CreateRedirection			godoc
// @Summary 					Creates a new redirection
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Param						request body models.RedirectionInput true "query params"
// @Success						200 	{array} 	models.Redirection
// @Failure						400 	{object} 	error
// @Router						/redirection [post]
func CreateRedirection(c *gin.Context) {
	// Input validation
	var input models.RedirectionInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": redirection})
}

// DELETE /redirection/:id

// DeleteRedirection			godoc
// @Summary 					Deletes a redirection in function of its ID
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Param						id 		path		int true "id"
// @Success						200 	{object} 	boolean
// @Failure						404 	{object} 	error
// @Router						/redirection/{id} [delete]
func DeleteRedirection(c *gin.Context) {
	// Gets model if exists
	var redirection models.Redirection
	err := database.DB.Where("id = ?", c.Param("id")).First(&redirection).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": errors.NotFoundError})
		return
	}

	database.DB.Delete(&redirection)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
