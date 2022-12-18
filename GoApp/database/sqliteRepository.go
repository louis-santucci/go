package database

import (
	"database/sql"
	"errors"
	"go.com/models"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row doesn't exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS redirections(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		redirect_url TEXT NOT NULL,
		views INTEGER NOT NULL,
		created_at TEXT NOT NULL
   );
	`

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(redirection models.Redirection) (*models.Redirection, error) {
	//FIXME To implement
}

func (r *SQLiteRepository) All() ([]models.Redirection, error) {
	//FIXME To implement
}

func (r *SQLiteRepository) GetById(id int64) (*models.Redirection, error) {
	//FIXME To implement
}

func (r *SQLiteRepository) GetByName(name string) (*models.Redirection, error) {
	//FIXME To implement
}

func (r *SQLiteRepository) Update(id int64, updated models.Redirection) (*models.Redirection, error) {
	//FIXME To implement
}

func (r *SQLiteRepository) Delete(id int64) error {
	//FIXME To implement
}
