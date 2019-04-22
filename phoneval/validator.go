package phoneval

import (
	"log"
	"strconv"
	"strings"
	"unicode"

	"gitlab.com/samuelfvlcastro/talkdesk-challenge-sc/x/rtrie"
)

const (
	zeroRune rune = 48
	nineRune rune = 57

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

// FindAreaCode matches the given Number with the longest existing area code prefix
func (v Validator) FindAreaCode(number Number) (int, error) {
	result, err := v.acTrie.Search(number.Clean)
	if err != nil {
		return 0, err
	}

	if result == "" {
		return 0, nil
	}

	ac, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return ac, err
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
		if err := acTrie.Insert(code); err != nil {
			log.Printf("invalid areacode, skipping [%s]", code)
		}
	}

	return acTrie
}
