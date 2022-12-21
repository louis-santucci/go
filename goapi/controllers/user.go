package controllers

import (
	"github.com/gin-gonic/gin"
	"louissantucci/goapi/database"
	"louissantucci/goapi/errors"
	"louissantucci/goapi/middlewares/jwt"
	"louissantucci/goapi/models"
	"net/http"
	"time"
)

// POST /user/register

// RegisterUser 				godoc
// @Summary						Creates new User in DB
// @Tags						user
// @Accept						json
// @Produce						json
// @Param						request body models.UserInput true "query params"
// @Success						200		{object}	models.User
// @Router						/user/register [post]
func RegisterUser(c *gin.Context) {
	// Input validation
	var input models.UserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// User creation
	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		CreatedAt: time.Now(),
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = database.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /user/edit

// EditUser		 				godoc
// @Summary						Edits user in DB
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags						user
// @Accept						json
// @Produce						json
// @Param						id 		path		int true "id"
// @Param						request body models.UserInput true "query params"
// @Success						200		{object}	models.User
// @Router						/user/edit/{id} [post]
func EditUser(c *gin.Context) {
	jwtToken, err := jwt.ExtractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	email, err := jwt.GetEmailFromToken(jwtToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	err = database.DB.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": errors.NotFoundError})
		return
	}

	if user.Email != email {
		c.JSON(http.StatusForbidden, gin.H{"error": errors.ForbiddenError})
		return
	}

	// Input validation
	var input models.UserInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedUser := models.User{
		ID:        user.ID,
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: time.Now(),
	}
	err = database.DB.Model(&user).Updates(updatedUser).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /user/login

// LoginUser 					godoc
// @Summary						Generates JWT token by verifying user credentials
// @Tags						user
// @Accept						json
// @Produce						json
// @Param						request body models.UserLogin true "query params"
// @Success						200		{object}	models.SignedResponse
// @Failure						403
// @Router						/user/login [post]
func LoginUser(c *gin.Context) {
	var loginRequest models.UserLogin
	var user models.User

	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Credentials check
	entry := database.DB.Where("email = ?", loginRequest.Email).First(&user)
	if entry.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": entry.Error.Error()})
		return
	}
	credentialsError := user.ComparePassword(loginRequest.Password)
	if credentialsError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.UnauthorizedError})
		return
	}
	tokenStr, err := jwt.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}
