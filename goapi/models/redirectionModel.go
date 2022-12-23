package models

import (
	"time"
)

type Redirection struct {
	ID          int64     `json:"id" gorm:"primarykey"`
	Shortcut    string    `json:"shortcut" gorm:"unique;not null"`
	RedirectUrl string    `json:"redirect_url" gorm:"not null"`
	Views       uint      `json:"views" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
	CreatorId   int64     `json:"creator_id" gorm:"not null"`
}

type RedirectionInput struct {
	Shortcut    string `json:"shortcut" binding:"required"`
	RedirectUrl string `json:"redirect_url" binding:"required"`
}
