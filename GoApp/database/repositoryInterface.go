package database

import "go.com/models"

type Repository interface {
	Migrate() error
	Create(redirection models.Redirection) (*models.Redirection, error)
	All() ([]models.Redirection, error)
	GetById(id int64) (*models.Redirection, error)
	GetByRedirection(redirection string) (*models.Redirection, error)
	Update(id int64, updated models.Redirection) (*models.Redirection, error)
	Delete(id int64) error
}
