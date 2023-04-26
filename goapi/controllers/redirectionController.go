package controllers

import (
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"louissantucci/goapi/database"
	error_constants "louissantucci/goapi/error-constants"
	"louissantucci/goapi/middlewares/jwt"
	"louissantucci/goapi/models"
	"louissantucci/goapi/responses"
	"louissantucci/goapi/services"
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
	redirections, err := services.GetRedirections()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

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
	idStr := c.Param("id")
	id, err := uuid2.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}
	redirection, err := services.GetRedirection(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewOKResponse(redirection))
}

// PUT /redirection/:id

// IncrementRedirectionView 	godoc
// @Summary 					Increments number of view for one redirection in function of its ID
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
	redirection.LastVisited = time.Now()

	var newHistoryEntry models.HistoryEntry
	// Get User ID
	header := c.GetHeader("Authorization")
	claim, err := jwt.GetClaimFromHeader(header)
	if err == nil {
		newHistoryEntry = models.HistoryEntry{
			VisitedAt:     redirection.LastVisited,
			RedirectionId: redirection.ID,
			UserId:        claim.Id,
		}
	} else {
		newHistoryEntry = models.HistoryEntry{
			VisitedAt:     redirection.LastVisited,
			RedirectionId: redirection.ID,
		}
	}

	err = database.DB.Model(&newHistoryEntry).Create(&newHistoryEntry).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	err = database.DB.Model(&redirection).Updates(redirection).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewOKResponse(redirection))
}

// PATCH /redirection/:id

// ResetRedirectionView 		godoc
// @Summary 					Resets number of view for one redirection in fucntion of its ID
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Param						id 		path		int true "id"
// @Success						200 	{object} 	responses.OKResponse
// @Failure						404 	{object} 	responses.ErrorResponse
// @Router						/redirection/{id} [patch]
func ResetRedirectionView(c *gin.Context) {
	var redirection models.Redirection

	err := database.DB.Where("id = ?", c.Param("id")).First(&redirection).Error
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	authHeader := c.GetHeader("Authorization")
	claim, err := jwt.GetClaimFromHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	if claim.Id != redirection.CreatorId {
		c.JSON(http.StatusUnauthorized, responses.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}

	// Reset
	var updateRedirection map[string]interface{}
	updateRedirection = map[string]interface{}{
		"Views": 0,
	}
	err = database.DB.Model(&redirection).Where("id = ?", redirection.ID).Updates(&updateRedirection).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

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
	// Input validation
	var input models.RedirectionInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}
	idStr := c.Param("id")
	id, err := uuid2.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	authHeader := c.GetHeader("Authorization")
	claim, err := jwt.GetClaimFromHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	redirection, err := services.EditRedirection(id, claim.Id, input)
	if err != nil {
		if err.Error() == error_constants.ForbiddenError {
			c.JSON(http.StatusUnauthorized, responses.NewErrorResponse(http.StatusUnauthorized, err.Error()))
			return
		}
		if err.Error() == error_constants.WrongRedirectionFormatError {
			c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
			return
		}
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

	header := c.GetHeader("Authorization")
	claim, err := jwt.GetClaimFromHeader(header)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	redirection, err := services.CreateRedirection(claim.Id, &input)
	if err != nil {
		switch err.Error() {
		case error_constants.WrongRedirectionFormatError:
			c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		case error_constants.NotFoundError:
			c.JSON(http.StatusNotFound, responses.NewErrorResponse(http.StatusNotFound, err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, responses.NewOKResponse(redirection))
}

// DELETE /redirection/:id

// DeleteRedirection			godoc
// @Summary 					Deletes a redirection in function of its ID
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags						redirection
// @Accept						json
// @Produce						json
// @Param						id 		path		int true "id"
// @Success						200 	{object} 	responses.OKResponse
// @Failure						404 	{object} 	responses.ErrorResponse
// @Failure						500 	{object} 	responses.ErrorResponse
// @Router						/redirection/{id} [delete]
func DeleteRedirection(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid2.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}
	authHeader := c.GetHeader("Authorization")
	claim, err := jwt.GetClaimFromHeader(authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	err = services.DeleteRedirection(id, claim.Id)
	if err != nil {
		switch err.Error() {
		case error_constants.UnauthorizedError:
			c.JSON(http.StatusUnauthorized, responses.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		case error_constants.NotFoundError:
			c.JSON(http.StatusNotFound, responses.NewErrorResponse(http.StatusNotFound, err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	responseData := "Redirection #" + idStr + " deleted"
	c.JSON(http.StatusOK, responses.NewOKResponse(responseData))
}
