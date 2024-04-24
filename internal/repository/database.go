package repository

import (
	"database/sql"
	"errors"
	"go-url-shortener/internal/entity"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseShortenerRepository struct {
	db *sql.DB
}

func (repo *DatabaseShortenerRepository) GetByHash(hash string) (*entity.ShortenedUrlEntity, error) {
	query := "SELECT hash, original_url, expiration_date FROM urlshortener WHERE hash = ?"
	row := repo.db.QueryRow(query, hash)

	var url entity.ShortenedUrlEntity

	err := row.Scan(&url.Hash, &url.OriginalUrl, &url.ExpirationDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &url, nil
}

func (repo *DatabaseShortenerRepository) GetByOriginalUrl(originalUrl string) (*entity.ShortenedUrlEntity, error) {
	query := "SELECT hash, original_url, expiration_date FROM urlshortener WHERE original_url = ?"
	row := repo.db.QueryRow(query, originalUrl)

	var url entity.ShortenedUrlEntity

	err := row.Scan(&url.Hash, &url.OriginalUrl, &url.ExpirationDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &url, nil
}

func (repo *DatabaseShortenerRepository) Save(url *entity.ShortenedUrlEntity) error {
	query := "INSERT INTO urlshortener (hash, original_url, expiration_date) VALUES (?, ?, ?)"

	_, err := repo.db.Exec(query, url.Hash, url.OriginalUrl, url.ExpirationDate)
	if err != nil {
		return err
	}

	return nil
}
