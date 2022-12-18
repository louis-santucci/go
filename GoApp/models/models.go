package models

import "time"

type Redirection struct {
	ID           int64     `json:"id"`
	SHORTCUT     string    `json:"shortcut"`
	REDIRECT_URL string    `json:"redirect_url"`
	VIEWS        int64     `json:"views"`
	CREATED_AT   time.Time `json:"created_at"`
}
