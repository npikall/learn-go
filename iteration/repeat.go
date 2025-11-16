package iteration

import "strings"

const repeatCount = 5

func RepeatSimple(character string) string {
	var repeated string
	for range repeatCount {
		repeated += character
	}
	return repeated
}

func RepeatImproved(character string) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(character)
	}
	return repeated.String()
}
