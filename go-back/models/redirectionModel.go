package models

import (
	"time"
)

type Redirection struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	Shortcut    string    `json:"shortcut" gorm:"unique;not null"`
	RedirectUrl string    `json:"redirect_url" gorm:"not null"'`
	Views       int64     `json:"views" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
}

type RedirectionCreationInput struct {
	Shortcut    string    `json:"shortcut"`
	RedirectUrl string    `json:"redirect_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type RedirectionUpdateInput struct {
	Shortcut    string    `json:"shortcut"`
	RedirectUrl string    `json:"redirect_url"`
	Views       int64     `json:"views"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RedirectionIncrement struct {
	Views int64 `json:"views"`
}
