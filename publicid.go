// Package publicid generates URL-safe, non-sequential public identifiers
// using a fixed nanoid alphabet [0-9a-z] and a default length of 24.
//
// PlanetScale reference for length sizing under collision constraints:
// https://planetscale.com/blog/why-we-chose-nanoids-for-planetscales-api
package publicid

import (
	"fmt"
	"strings"

	nanoid "github.com/matoous/go-nanoid/v2"
)

// Alphabet is the lowercase alphanumeric character set used by every public ID.
const Alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"

// DefaultLength is the default character length for IDs returned by New / Must.
const DefaultLength = 24

// New generates a public ID of the default length.
func New() (string, error) {
	return nanoid.Generate(Alphabet, DefaultLength)
}

// Must is the same as New but panics on error.
//
// Errors from nanoid.Generate are essentially impossible (math/rand never errors,
// crypto/rand only fails on broken systems), so Must is the convention for
// hot-path call sites that don't want the err return.
func Must() string {
	return nanoid.MustGenerate(Alphabet, DefaultLength)
}

// NewN generates a public ID of length n.
func NewN(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("publicid: NewN length must be > 0, got %d", n)
	}
	return nanoid.Generate(Alphabet, n)
}

// MustN is the same as NewN but panics on error.
func MustN(n int) string {
	if n <= 0 {
		panic(fmt.Sprintf("publicid: MustN length must be > 0, got %d", n))
	}
	return nanoid.MustGenerate(Alphabet, n)
}

// Validate checks that id is a non-empty string of the default length and
// contains only characters from Alphabet. fieldName is included in error
// messages for easier debugging.
func Validate(fieldName, id string) error {
	return ValidateN(fieldName, id, DefaultLength)
}

// ValidateN is the same as Validate but checks for a custom expected length.
func ValidateN(fieldName, id string, n int) error {
	if id == "" {
		return fmt.Errorf("%s cannot be blank", fieldName)
	}
	if len(id) != n {
		return fmt.Errorf("%s should be %d characters long", fieldName, n)
	}
	if strings.Trim(id, Alphabet) != "" {
		return fmt.Errorf("%s has invalid characters", fieldName)
	}
	return nil
}
