package phoneval

import "strings"

// AreaCode prefix used to differentiate phone area codes
type AreaCode string

// Number phone number data and localization
type Number struct {
	Area     AreaCode
	Clean    string
	Original string
}

// New factory method to instantiate a phone number and generate a standardize number
func New(code AreaCode, number string) Number {
	clean := standardizeNumber(number)
	return Number{
		Area:     code,
		Clean:    clean,
		Original: number,
	}
}

func standardizeNumber(number string) string {
	return strings.TrimLeft(strings.Replace(number, " ", "", -1), "+00")
}
