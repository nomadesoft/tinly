// Package shortener contains implementation of URL shortener algorithm.
package shortener

// Shortener interface defines the shrink method.
type Shortener interface {

	// Shrink returns a shortURL from original URL.
	Shrink(originalURL string) (shortURL string, err error)
}
