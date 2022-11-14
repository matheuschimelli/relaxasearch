package utils

import (
	"log"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func Normalize(str string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, str)

	re, err := regexp.Compile(`[^\w]`)

	if err != nil {
		log.Fatal(err)
	}

	finalResult := re.ReplaceAllString(result, " ")
	finalResult = strings.ToLower(finalResult)

	return finalResult
}
