package database

import (
	"database/sql"
	"errors"
	"github.com/mattn/go-sqlite3"
	"go.com/models"
	"time"
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
	_, err := r.db.Exec(migrateQuery)
	return err
}

func (r *SQLiteRepository) Create(redirection models.Redirection) (*models.Redirection, error) {
	result, err := r.db.Exec(createQuery, redirection.SHORTCUT, redirection.REDIRECT_URL, 0, time.Now().String())
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	redirection.ID = id

	return &redirection, nil
}

func (r *SQLiteRepository) All() ([]models.Redirection, error) {
	results, err := r.db.Query(selectAllQuery)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var redirections []models.Redirection
	for results.Next() {
		var redirection models.Redirection
		err := results.Scan(&redirection.ID, &redirection.SHORTCUT, &redirection.REDIRECT_URL, &redirection.VIEWS, &redirection.CREATED_AT)
		if err != nil {
			return nil, err
		}
		redirections = append(redirections, redirection)
	}
	return redirections, nil
}

func (r *SQLiteRepository) GetById(id int64) (*models.Redirection, error) {
	row := r.db.QueryRow(selectByIdQuery, id)
	var redirection models.Redirection
	err := row.Scan(&redirection.ID, &redirection.SHORTCUT, &redirection.REDIRECT_URL, &redirection.VIEWS, &redirection.CREATED_AT)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &redirection, nil
}

func (r *SQLiteRepository) GetByRedirection(redirection string) (*models.Redirection, error) {
	row := r.db.QueryRow(selectByRedirectionQuery, redirection)
	var redirectionModel models.Redirection
	err := row.Scan(&redirectionModel.ID, &redirectionModel.SHORTCUT, &redirectionModel.REDIRECT_URL, &redirectionModel.VIEWS, &redirectionModel.CREATED_AT)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}

	return &redirectionModel, nil
}

func (r *SQLiteRepository) Update(id int64, updated models.Redirection) (*models.Redirection, error) {
	if id <= 0 {
		return nil, ErrInvalidId
	}

	res, err := r.db.Exec(updateQuery, updated.SHORTCUT, updated.REDIRECT_URL, updated.VIEWS, updated.CREATED_AT)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}

func (r *SQLiteRepository) Delete(id int64) error {
	if id <= 0 {
		return ErrInvalidId
	}
	res, err := r.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
