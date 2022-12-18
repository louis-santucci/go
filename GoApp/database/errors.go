package database

import "errors"

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row doesn't exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
	ErrInvalidId    = errors.New("invalid given id")
)
