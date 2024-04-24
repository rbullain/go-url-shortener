package service

import (
	"go-url-shortener/internal/entity"
	"go-url-shortener/internal/repository"
	"go-url-shortener/internal/utils"
	"time"
)

type ShortenerService interface {
	ShortenUrl(url string) (*entity.ShortenedUrlEntity, error)
	ExpandUrl(url string) (*entity.ShortenedUrlEntity, error)
}

type shortenerService struct {
	repository repository.ShortenerRepository
}

func NewShortenerService(repository repository.ShortenerRepository) ShortenerService {
	return &shortenerService{
		repository: repository,
	}
}

func (service *shortenerService) ShortenUrl(url string) (*entity.ShortenedUrlEntity, error) {
	existingUrl, err := service.repository.GetByOriginalUrl(url)
	if err != nil {
		return nil, err
	}
	if existingUrl != nil {
		return existingUrl, nil
	}

	newUrl := &entity.ShortenedUrlEntity{
		Hash:           utils.GenerateHash(url),
		OriginalUrl:    url,
		ExpirationDate: time.Now().AddDate(0, 0, 1),
	}
	err = service.repository.Save(newUrl)
	if err != nil {
		return nil, err
	}

	return newUrl, nil
}

func (service *shortenerService) ExpandUrl(url string) (*entity.ShortenedUrlEntity, error) {
	existingUrl, err := service.repository.GetByHash(url)
	if err != nil {
		return nil, err
	}
	return existingUrl, nil
}
