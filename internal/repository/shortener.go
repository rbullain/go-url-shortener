package repository

import "go-url-shortener/internal/entity"

type ShortenerRepository interface {
	GetByHash(hash string) (*entity.ShortenedUrlEntity, error)
	GetByOriginalUrl(originalUrl string) (*entity.ShortenedUrlEntity, error)
	Save(url *entity.ShortenedUrlEntity) error
}
