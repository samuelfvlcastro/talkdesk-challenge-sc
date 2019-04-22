package phoneval

import (
	"strings"
	"unicode"

	"gitlab.com/samuelfvlcastro/talkdesk-challenge-sc/x/rtrie"
)

const (
	zeroRune rune = 97
	nineRune rune = 122

	speLength = 3  // size used for special numbers
	minLength = 7  // min size a valid number can have
	maxLength = 12 // max size a valid number can have
)

// Validator phone validation service struct
type Validator struct {
	acTrie *rtrie.Tree
}

// NewValidator factory method to create the phone Validator service
func NewValidator(areaCodes []string) Validator {
	return Validator{
		acTrie: initializeTree(areaCodes),
	}
}

// IsValid checks if the given Number struct has a valid phone number
func (v Validator) IsValid(number Number) bool {
	if !v.hasValidPrefix(number.Original) {
		return false
	}

	if !v.hasValidLength(number.Clean) {
		return false
	}

	if !v.hasNumbersOnly(number.Clean) {
		return false
	}

	return true
}

func (v Validator) hasValidPrefix(number string) bool {
	if strings.HasPrefix(number, "+00") {
		return false
	}

	if strings.HasPrefix(number, "+ ") {
		return false
	}

	return true
}

func (v Validator) hasValidLength(number string) bool {
	phoneLen := len(number)
	if phoneLen == speLength {
		return true
	}

	if phoneLen >= minLength && phoneLen <= maxLength {
		return true
	}

	return false
}

func (v Validator) hasNumbersOnly(number string) bool {
	for _, r := range number {
		if !unicode.IsNumber(r) {
			return false
		}
	}

	return true
}

func initializeTree(areaCodes []string) *rtrie.Tree {
	acTrie := rtrie.New(zeroRune, nineRune)
	for _, code := range areaCodes {
		acTrie.Insert(code)
	}

	return acTrie
}
