package main

import (
	"go.com/models"
	"time"
)

var redirections = []models.Redirection{
	{ID: 1, SHORTCUT: "fb", REDIRECT_URL: "https://facebook.com", CREATED_AT: time.Now()},
	{ID: 2, SHORTCUT: "fb", REDIRECT_URL: "https://facebook.com", CREATED_AT: time.Now()},
	{ID: 3, SHORTCUT: "fb", REDIRECT_URL: "https://facebook.com", CREATED_AT: time.Now()},
	{ID: 4, SHORTCUT: "fb", REDIRECT_URL: "https://facebook.com", CREATED_AT: time.Now()},
	{ID: 5, SHORTCUT: "fb", REDIRECT_URL: "https://facebook.com", CREATED_AT: time.Now()},
}
