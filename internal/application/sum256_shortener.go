package application

import (
	"crypto/sha256"
	"encoding/base64"
	"go-url-shortener/internal/domain/shortener/entity"
	"go-url-shortener/internal/domain/shortener/service"
	"time"
)

type sum256HashShortener struct {
	repository entity.ShortenerRepository
}

func NewSum256HashShortener(repository entity.ShortenerRepository) service.ShortenerService {
	return &sum256HashShortener{
		repository: repository,
	}
}

func generateHash(s string) string {
	hash := sha256.Sum256([]byte(s))
	encodedHash := base64.URLEncoding.EncodeToString(hash[:])
	return encodedHash[:8]
}

func (service *sum256HashShortener) ShortenUrl(url string) (*entity.ShortenedUrlEntity, error) {
	existingUrl, err := service.repository.GetByOriginalUrl(url)
	if err != nil {
		return nil, err
	}
	if existingUrl != nil {
		return existingUrl, nil
	}

	newUrl := &entity.ShortenedUrlEntity{
		Hash:           generateHash(url),
		OriginalUrl:    url,
		ExpirationDate: time.Now().AddDate(0, 0, 1),
	}
	err = service.repository.Save(newUrl)
	if err != nil {
		return nil, err
	}

	return newUrl, nil
}

func (service *sum256HashShortener) ExpandUrl(url string) (*entity.ShortenedUrlEntity, error) {
	existingUrl, err := service.repository.GetByHash(url)
	if err != nil {
		return nil, err
	}
	return existingUrl, nil
}
