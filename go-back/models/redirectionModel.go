package models

import (
	"time"
)

type Redirection struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	Shortcut    string    `json:"shortcut" gorm:"unique;not null"`
	RedirectUrl string    `json:"redirect_url" gorm:"not null"`
	Views       int64     `json:"views" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
}

type RedirectionInput struct {
	Shortcut    string `json:"shortcut" binding:"required"`
	RedirectUrl string `json:"redirect_url" binding:"required"`
}

type RedirectionIncrement struct {
	Views int64 `json:"views"`
}
