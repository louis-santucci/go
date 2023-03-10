package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type HistoryEntry struct {
	ID            uuid.UUID `json:"id" gorm:"primarykey"`
	RedirectionId uuid.UUID `json:"redirection_id" gorm:"not null"`
	VisitedAt     time.Time `json:"visited_at" gorm:"not null"`
	UserId        uuid.UUID `json:"user_id"`
}

func (historyEntry *HistoryEntry) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	historyEntry.ID = uuid.New()
	return
}
