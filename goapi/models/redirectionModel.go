package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Redirection struct {
	ID          uuid.UUID `json:"id" gorm:"primarykey"`
	Shortcut    string    `json:"shortcut" gorm:"unique;not null"`
	RedirectUrl string    `json:"redirect_url" gorm:"not null"`
	Views       uint      `json:"views" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at;not null"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
	LastVisited time.Time `json:"last_visited"`
	CreatorId   uuid.UUID `json:"creator_id" gorm:"not null"`
}

type RedirectionInput struct {
	Shortcut    string `json:"shortcut" binding:"required"`
	RedirectUrl string `json:"redirect_url" binding:"required"`
}

func (redirection *Redirection) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	redirection.ID = uuid.New()
	return
}
