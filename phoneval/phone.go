package phoneval

import "strings"

// Number phone number data and localization
type Number struct {
	Clean    string
	Original string
}

// NewNumber factory method to instantiate a phone number and generate a standardize number
func NewNumber(number string) Number {
	clean := standardizeNumber(number)
	return Number{
		Clean:    clean,
		Original: number,
	}
}

func standardizeNumber(number string) string {
	return strings.TrimPrefix(strings.TrimPrefix(strings.Replace(number, " ", "", -1), "+"), "00")
}
