package shared

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

func SlugString(term string) string {
	lowerCase := cases.Lower(language.English, cases.NoLower)
	lowerTerm := lowerCase.String(term)

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	cleaned, _, _ := transform.String(t, lowerTerm)
	cleaned = strings.ReplaceAll(cleaned, " ", "-")
	cleaned = strings.ReplaceAll(cleaned, "_", "-")

	return cleaned
}
