package slug

import (
	"errors"
	"regexp"
	"strings"

	apperrors "github.com/onurhan1337/slugo/pkg/errors"
)

func Generate(text string) (string, error) {
	s := strings.ToLower(text)

	s = strings.ReplaceAll(s, "ç", "c")
	s = strings.ReplaceAll(s, "ğ", "g")
	s = strings.ReplaceAll(s, "ı", "i")
	s = strings.ReplaceAll(s, "ö", "o")
	s = strings.ReplaceAll(s, "ş", "s")
	s = strings.ReplaceAll(s, "ü", "u")

	regNonAlphaNumeric := regexp.MustCompile(`[^a-z0-9\s-]`)
	s = regNonAlphaNumeric.ReplaceAllString(s, " ")

	regMultipleSpaces := regexp.MustCompile(`\s+`)
	s = regMultipleSpaces.ReplaceAllString(s, " ")

	trimmedS := strings.TrimSpace(s)

	if trimmedS == "" {
		var reason string
		if strings.TrimSpace(text) == "" {
			reason = "text consists only of whitespace"
		} else {
			reason = "text contains only invalid characters"
		}
		return "", apperrors.NewValidationError("input", text, errors.New(reason))
	}

	s = strings.ReplaceAll(trimmedS, " ", "-")

	regMultipleDashes := regexp.MustCompile(`-+`)
	s = regMultipleDashes.ReplaceAllString(s, "-")

	return s, nil
}
