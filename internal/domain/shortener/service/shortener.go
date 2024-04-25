package service

import "go-url-shortener/internal/domain/shortener/entity"

type ShortenerService interface {
	ShortenUrl(url string) (*entity.ShortenedUrlEntity, error)
	ExpandUrl(url string) (*entity.ShortenedUrlEntity, error)
}
