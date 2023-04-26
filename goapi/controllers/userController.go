package controllers

import (
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"louissantucci/goapi/database"
	"louissantucci/goapi/error-constants"
	"louissantucci/goapi/middlewares/jwt"
	"louissantucci/goapi/models"
	"louissantucci/goapi/responses"
	"louissantucci/goapi/services"
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

	currentTime := time.Now()
	// User creation
	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
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

// GET /user/info

// GetUserInfo		 				godoc
// @Summary						Returns User information with Auth header
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags						user
// @Accept						json
// @Produce						json
// @Success						200 	{object} 	responses.OKResponse
// @Failure						500 	{object} 	responses.ErrorResponse
// @Router						/user/info [get]
func GetUserInfo(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	claim, err := jwt.GetClaimFromToken(authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	userInfo := services.GetUserInfo(claim.Id)

	c.JSON(http.StatusOK, responses.NewOKResponse(userInfo))
}

// DELETE /user/delete

// DeleteUser	 				godoc
// @Summary						Deletes User with Auth header
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags						user
// @Accept						json
// @Produce						json
// @Success						200 	{object} 	responses.OKResponse
// @Failure						500 	{object} 	responses.ErrorResponse
// @Router						/user/delete [delete]
func DeleteUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	jwtToken, err := jwt.ExtractBearerToken(authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	claim, err := jwt.GetClaimFromToken(jwtToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	err = services.DeleteUser(claim.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	message := "User deleted"
	c.JSON(http.StatusOK, responses.NewOKResponse(message))
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
	authHeader := c.GetHeader("Authorization")
	claim, err := jwt.GetClaimFromToken(authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	idStr := c.Param("id")
	id, err := uuid2.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// Input validation
	var input models.UserInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	updatedUser, err := services.EditUser(id, claim.Id, input)
	if err != nil {
		switch err.Error() {
		case error_constants.UnauthorizedError:
			c.JSON(http.StatusUnauthorized, responses.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, responses.NewOKResponse(updatedUser))
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
// @Failure						404 	{object}  	responses.ErrorResponse
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

	tokenStr, err := services.LoginUser(loginRequest)
	if err != nil {
		switch err.Error() {
		case error_constants.UnauthorizedError:
			c.JSON(http.StatusUnauthorized, responses.NewErrorResponse(http.StatusUnauthorized, err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return
	}
	c.JSON(http.StatusOK, responses.NewJWTResponse(http.StatusOK, *tokenStr, user.Email))
}

// GetUsers			godoc
// @Summary 					Get all users
// @Tags						user
// @Accept						json
// @Produce						json
// @Success						200 	{object} 	responses.OKResponse
// @Router						/user/list [get]
func GetUsers(c *gin.Context) {
	userInfos, err := services.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, responses.NewOKResponse(userInfos))
}
