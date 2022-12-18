package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) ComparePassword(givenPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(givenPassword))
	if err != nil {
		return err
	}
	return nil
}
