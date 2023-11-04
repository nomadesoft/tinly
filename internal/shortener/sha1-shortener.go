package shortener

import (
	"crypto/sha1"
	"encoding/hex"
	"hash"
)

type Sha1Shortener struct {
	hasher hash.Hash
}

func NewSha1Shortener() *Sha1Shortener {
	return &Sha1Shortener{
		hasher: sha1.New(),
	}
}

func (s *Sha1Shortener) Shrink(originalURL string) (shortURL string, err error) {
	// Calculate hash
	s.hasher.Write([]byte(originalURL))
	hashBytes := s.hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	// Trunk hash to 7 chars
	shortURL = hashString[:7]

	return shortURL, nil
}
