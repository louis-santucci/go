package models

import "time"

type Redirection struct {
	ID           float64   `json:"id"`
	SHORTCUT     string    `json:"shortcut"`
	REDIRECT_URL string    `json:"redirect_url"`
	CREATED_AT   time.Time `json:"created_at"`
}
