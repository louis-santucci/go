package controllers

import (
	"github.com/gin-gonic/gin"
	"louissantucci/goapi/database"
	"louissantucci/goapi/errors"
	"louissantucci/goapi/middlewares/jwt"
	"louissantucci/goapi/models"
	"louissantucci/goapi/responses"
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
// @Success						200 	{object} 	responses.OKResponse
// @Failure						400 	{object} 	responses.ErrorResponse
// @Failure						500 	{object} 	responses.ErrorResponse
// @Router						/user/register [post]
func RegisterUser(c *gin.Context) {
	// Input validation
	var input models.UserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
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
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	err = database.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.NewOKResponse(user))
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
// @Success						200 	{object} 	responses.OKResponse
// @Failure						400 	{object} 	responses.ErrorResponse
// @Failure						403 	{object} 	responses.ErrorResponse
// @Failure						404 	{object} 	responses.ErrorResponse
// @Failure						500 	{object} 	responses.ErrorResponse
// @Router						/user/edit/{id} [post]
func EditUser(c *gin.Context) {
	jwtToken, err := jwt.ExtractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	email, err := jwt.GetEmailFromToken(jwtToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	var user models.User

	err = database.DB.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NewErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	if user.Email != email {
		c.JSON(http.StatusForbidden, responses.NewErrorResponse(http.StatusForbidden, err.Error()))
		return
	}

	// Input validation
	var input models.UserInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	user.UpdatedAt = time.Now()
	err = database.DB.Model(&user).Updates(user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewOKResponse(user))
}

// POST /user/login

// LoginUser 					godoc
// @Summary						Generates JWT token by verifying user credentials
// @Tags						user
// @Accept						json
// @Produce						json
// @Param						request body models.UserLogin true "query params"
// @Success						200		{object}	responses.JWTResponse
// @Failure						400 	{object}  	responses.ErrorResponse
// @Failure						401 	{object}  	responses.ErrorResponse
// @Failure						500 	{object}  	responses.ErrorResponse
// @Router						/user/login [post]
func LoginUser(c *gin.Context) {
	var loginRequest models.UserLogin
	var user models.User

	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// Credentials check
	entry := database.DB.Where("email = ?", loginRequest.Email).First(&user)
	if entry.Error != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, entry.Error.Error()))
		return
	}
	credentialsError := user.ComparePassword(loginRequest.Password)
	if credentialsError != nil {
		errorData := errors.UnauthorizedError + ": " + credentialsError.Error()
		c.JSON(http.StatusUnauthorized, responses.NewErrorResponse(http.StatusUnauthorized, errorData))
		return
	}
	tokenStr, err := jwt.GenerateJWT(user.Email, user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.NewJWTResponse(http.StatusOK, tokenStr))
}
