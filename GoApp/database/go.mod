module go.com/database

go 1.19

// Replaces go.com/models package by current local one i.e. ../models to fix import error
replace go.com/models => ../models

require (
	github.com/mattn/go-sqlite3 v1.14.16
	go.com/models v0.0.0-00010101000000-000000000000
)
