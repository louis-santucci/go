package services

import (
	"github.com/google/uuid"
	"louissantucci/goapi/database"
	"louissantucci/goapi/models"
)

func DeleteRedirectionHistory(id uuid.UUID) error {
	return database.DB.Delete(&models.HistoryEntry{}, "redirection_id = ?", id).Error
}

func ResetHistory() error {
	return database.DB.Exec("DELETE FROM history_entries").Error
}

func GetHistory() ([]models.HistoryEntry, error) {
	var historyResults []models.HistoryEntry
	err := database.DB.Find(&historyResults).Error
	return historyResults, err
}
