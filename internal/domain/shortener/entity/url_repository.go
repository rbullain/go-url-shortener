package entity

type ShortenerRepository interface {
	GetByHash(hash string) (*ShortenedUrlEntity, error)
	GetByOriginalUrl(originalUrl string) (*ShortenedUrlEntity, error)
	Save(url *ShortenedUrlEntity) error
}
