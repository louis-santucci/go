package services

import (
	"errors"
	"github.com/google/uuid"
	"louissantucci/goapi/database"
	error_constants "louissantucci/goapi/error-constants"
	"louissantucci/goapi/models"
	"regexp"
	"time"
)

const URL_REGEX = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`

func GetRedirection(id uuid.UUID) (*models.Redirection, error) {
	var redirection models.Redirection

	err := database.DB.Where("id = ?", id).First(&redirection).Error
	if err != nil {
		return nil, err
	}

	return &redirection, nil
}

func GetRedirections() ([]models.Redirection, error) {
	var redirections []models.Redirection
	err := database.DB.Find(&redirections).Error
	return redirections, err
}

func CreateRedirection(userId uuid.UUID, input *models.RedirectionInput) (*models.Redirection, error) {
	// Check redirection format
	if !ValidateRedirectUrlFormat(input.RedirectUrl) {
		return nil, errors.New(error_constants.WrongRedirectionFormatError)
	}

	user, err := GetUser(userId)
	if err != nil {
		return nil, errors.New(error_constants.NotFoundError)
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
		return nil, err
	}
	return &redirection, nil
}

func EditRedirection(id uuid.UUID, userId uuid.UUID, input models.RedirectionInput) (*models.Redirection, error) {
	redirection, err := GetRedirection(id)
	if err != nil {
		return nil, err
	}

	if userId != redirection.CreatorId {
		return nil, errors.New(error_constants.ForbiddenError)
	}

	// Check redirection format
	if !ValidateRedirectUrlFormat(input.RedirectUrl) {
		return nil, errors.New(error_constants.WrongRedirectionFormatError)
	}

	redirection.Shortcut = input.Shortcut
	redirection.RedirectUrl = input.RedirectUrl
	redirection.UpdatedAt = time.Now()
	err = database.DB.Model(&redirection).Updates(redirection).Error
	if err != nil {
		return nil, err
	}

	return redirection, nil
}

func DeleteRedirection(id uuid.UUID, userId uuid.UUID) error {
	// Gets model if exists
	var redirection models.Redirection
	err := database.DB.Where("id = ?", id).First(&redirection).Error
	if err != nil {
		return errors.New(error_constants.NotFoundError)
	}
	// Checks if user is creator
	if userId != redirection.CreatorId {
		return errors.New(error_constants.UnauthorizedError)
	}
	// Deletes redirection
	err = database.DB.Delete(&redirection).Error
	if err != nil {
		return err
	}
	// Deletes all history entries for the redirection
	return DeleteRedirectionHistory(id)
}

func ValidateRedirectUrlFormat(url string) bool {
	regex, _ := regexp.Compile(URL_REGEX)
	return regex.MatchString(url)
}
