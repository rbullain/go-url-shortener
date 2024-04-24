package entity

import "time"

type ShortenedUrlEntity struct {
	Hash           string    `json:"hash"`
	OriginalUrl    string    `json:"original_url"`
	ExpirationDate time.Time `json:"expiration_date"`
}
