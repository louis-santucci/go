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

func EditUser(userId uuid.UUID, input models.UserInput) (*models.User, error) {
	user, err := GetUser(userId)
	if err != nil {
		return nil, err
	}
	credentialsError := user.ComparePassword(input.Password)
	if credentialsError != nil {
		return nil, errors.New(error_constants.UnauthorizedError)
	}

	user.Name = input.Name
	user.Email = input.Email
	user.UpdatedAt = time.Now()

	err = database.DB.Model(user).Updates(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func LoginUser(email string, password string) (*models.UserInfo, *string, error) {
	var user models.User
	// Credentials check
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, nil, err
	}
	credentialsError := user.ComparePassword(password)
	if credentialsError != nil {
		return nil, nil, errors.New(error_constants.UnauthorizedError)
	}
	tokenStr, err := jwt.GenerateJWT(&user)
	if err != nil {
		return nil, nil, err
	}

	return GetUserInfo(user.ID), &tokenStr, nil
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
