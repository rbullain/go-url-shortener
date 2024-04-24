package dto

import "time"

type ShortenUrlRequestDTO struct {
	Url string `json:"url"`
}

type ShortenUrlResponseDTO struct {
	Url            string    `json:"url"`
	ShortUrl       string    `json:"short_url"`
	ExpirationDate time.Time `json:"expiration_date"`
}
