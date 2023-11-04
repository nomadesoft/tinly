package shortener

import (
	"crypto/sha1"
	"encoding/hex"
)

type Sha1Shortener struct {
}

func (s *Sha1Shortener) Shrink(originalURL string) (shortURL string, err error) {
	hasher := sha1.New()

	hasher.Write([]byte(originalURL))

	hashBytes := hasher.Sum(nil)

	hashString := hex.EncodeToString(hashBytes)

	shortURL = hashString[:7]

	return shortURL, nil
}

func NewSha1Shortener() *Sha1Shortener {
	return &Sha1Shortener{}
}
