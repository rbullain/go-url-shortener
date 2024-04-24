package entity

import "time"

type ShortenedUrlEntity struct {
	Hash           string
	OriginalUrl    string
	ExpirationDate time.Time
}
