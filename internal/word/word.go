package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func UnderscoreToUpperCamelCase(s string) string {
	// replace all _ to space, -1 means do replace unlimited times.
	s = strings.Replace(s, "_", " ", -1)

	// convert all words separated by space to Capitalize.
	s = strings.Title(s)

	// merge all words by remove their space.
	return strings.Replace(s, " ", "", -1)
}

func UnderscoreToLowerCamelCase(s string) string {
	// first convert to CamelCase.
	s = UnderscoreToUpperCamelCase(s)

	// directly modify the first char of s to Lowercase, then concat with leftover chars.
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func CamelCaseToUnderScore(s string) string {
	var result []rune

	for i, r := range s {
		// the first word should always be convert to lowercase, ex: myApple MyApple => 'm'y_apple.
		if i == 0 {
			result = append(result, unicode.ToLower(r))
			continue
		}

		// if find a uppercase char, then convert it from 'A' to '_a'.
		if unicode.IsUpper(r) {
			result = append(result, '_')
			result = append(result, unicode.ToLower(r))
			continue
		}

		// lowercase word directly append to result.
		result = append(result, r)
	}

	return string(result)
}
