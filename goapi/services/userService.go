package services

import (
	"errors"
	"github.com/google/uuid"
	"louissantucci/goapi/database"
	error_constants "louissantucci/goapi/error-constants"
	"louissantucci/goapi/middlewares/jwt"
	"louissantucci/goapi/models"
	"time"
)

func GetUsers() ([]models.UserInfo, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	userInfos := make([]models.UserInfo, len(users))
	for i := 0; i < len(users); i++ {
		// User creation
		userInfo := models.UserInfo{
			ID:        users[i].ID,
			Name:      users[i].Name,
			Email:     users[i].Email,
			CreatedAt: users[i].CreatedAt,
			UpdatedAt: users[i].UpdatedAt,
		}
		userInfos[i] = userInfo
	}

	return userInfos, err
}

func DeleteUser(id uuid.UUID) error {
	return database.DB.Exec("DELETE FROM users WHERE id = ?", id).Error
}

func EditUser(id uuid.UUID, userId uuid.UUID, input models.UserInput) (*models.User, error) {
	user, err := GetUser(id)
	if err != nil {
		return nil, err
	}
	if userId != user.ID {
		return nil, errors.New(error_constants.ForbiddenError)
	}

	user.Name = input.Name
	user.Email = input.Email
	user.UpdatedAt = time.Now()

	err = user.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	err = database.DB.Model(user).Updates(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func LoginUser(loginRequest models.UserLogin) (*string, error) {
	var user models.User
	// Credentials check
	err := database.DB.Where("email = ?", loginRequest.Email).First(&user).Error
	if err != nil {
		return nil, err
	}
	credentialsError := user.ComparePassword(loginRequest.Password)
	if credentialsError != nil {
		return nil, errors.New(error_constants.UnauthorizedError)
	}
	tokenStr, err := jwt.GenerateJWT(&user)
	if err != nil {
		return nil, err
	}

	return &tokenStr, nil
}

func GetUserInfo(userId uuid.UUID) *models.UserInfo {
	user, err := GetUser(userId)
	if err != nil {
		return nil
	}

	// UserInfo creation
	return &models.UserInfo{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func GetUser(id uuid.UUID) (*models.User, error) {
	var user models.User

	err := database.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
