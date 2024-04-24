package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-url-shortener/internal/entity"

	_ "github.com/go-sql-driver/mysql"
)

const (
	getByHashQuery        = "SELECT hash, original_url, expiration_date FROM urlshortener WHERE hash = ?"
	getByOriginalUrlQuery = "SELECT hash, original_url, expiration_date FROM urlshortener WHERE original_url = ?"
	createUrlQuery        = "INSERT INTO urlshortener (hash, original_url, expiration_date) VALUES (?, ?, ?)"
)

type DatabaseShortenerRepository struct {
	db *sql.DB
}

func connect(username, password, host, port, database string) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewDatabaseShortenerRepository(username, password, host, port, database string) *DatabaseShortenerRepository {
	db, err := connect(username, password, host, port, database)
	if err != nil {
		panic(err)
	}

	return &DatabaseShortenerRepository{
		db: db,
	}
}

func (repo *DatabaseShortenerRepository) GetByHash(hash string) (*entity.ShortenedUrlEntity, error) {
	row := repo.db.QueryRow(getByHashQuery, hash)

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
	row := repo.db.QueryRow(getByOriginalUrlQuery, originalUrl)

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
	_, err := repo.db.Exec(createUrlQuery, url.Hash, url.OriginalUrl, url.ExpirationDate)
	if err != nil {
		return err
	}

	return nil
}
