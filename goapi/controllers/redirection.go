package controllers

import (
	"github.com/gin-gonic/gin"
	"louissantucci/goapi/database"
	"louissantucci/goapi/middlewares/jwt"
	"louissantucci/goapi/models"
	"louissantucci/goapi/responses"
	"net/http"
	"regexp"
	"time"
)

const URL_REGEX = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`

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
	var isLogged = false
	// Get User ID
	header := c.GetHeader("Authorization")
	tokenStr, err := jwt.ExtractBearerToken(header)
	if err == nil {
		email, err := jwt.GetEmailFromToken(tokenStr)
		if err == nil {
			var user models.User

			err = database.DB.Where("email = ?", email).First(&user).Error
			if err == nil {
				newHistoryEntry = models.HistoryEntry{
					VisitedAt:     redirection.LastVisited,
					RedirectionId: redirection.ID,
					UserId:        user.ID,
				}
				isLogged = true
			} else {
				isLogged = false
			}
		}
	}

	if !isLogged {
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
	errCode, err, _ := jwt.IsIdMatchingJwtToken(redirection.ID, authHeader)
	if err != nil {
		c.JSON(errCode, responses.NewErrorResponse(errCode, err.Error()))
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
	var redirection models.Redirection

	err := database.DB.Where("id = ?", c.Param("id")).First(&redirection).Error
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	authHeader := c.GetHeader("Authorization")
	errCode, err, _ := jwt.IsIdMatchingJwtToken(redirection.CreatorId, authHeader)
	if err != nil {
		c.JSON(errCode, responses.NewErrorResponse(errCode, err.Error()))
		return
	}

	// Input validation
	var input models.RedirectionInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// Check redirection format
	if !ValidateRedirectUrlFormat(input.RedirectUrl) {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, "Redirection URL must be a valid URL beginning with http[s]://"))
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

	// Get User ID
	header := c.GetHeader("Authorization")
	tokenStr, err := jwt.ExtractBearerToken(header)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	email, err := jwt.GetEmailFromToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	var user models.User

	err = database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// Check redirection format
	if !ValidateRedirectUrlFormat(input.RedirectUrl) {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, "Redirection URL must be a valid URL beginning with http[s]://"))
		return
	}

	// Redirection creation
	redirection := models.Redirection{
		Shortcut:    input.Shortcut,
		RedirectUrl: input.RedirectUrl,
		Views:       0,
		CreatedAt:   time.Now(),
		CreatorId:   user.ID,
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
	// Gets model if exists
	var redirection models.Redirection
	id := c.Param("id")
	err := database.DB.Where("id = ?", id).First(&redirection).Error
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	authHeader := c.GetHeader("Authorization")
	errCode, err, _ := jwt.IsIdMatchingJwtToken(redirection.CreatorId, authHeader)
	if err != nil {
		c.JSON(errCode, responses.NewErrorResponse(errCode, err.Error()))
		return
	}

	// Deletion
	err = database.DB.Delete(&redirection).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// Delete all history entries for this redirection
	err = database.DB.Delete(&models.HistoryEntry{}, "redirection_id = ?", redirection.ID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	responseData := "Redirection #" + id + " deleted"
	c.JSON(http.StatusOK, responses.NewOKResponse(responseData))
}

func ValidateRedirectUrlFormat(url string) bool {
	regex, _ := regexp.Compile(URL_REGEX)
	return regex.MatchString(url)
}
