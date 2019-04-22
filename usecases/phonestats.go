package usecases

import (
	"log"
	"sort"

	"gitlab.com/samuelfvlcastro/talkdesk-challenge-sc/phoneval"
)

type Validator interface {
	IsValid(number phoneval.Number) bool
	FindAreaCode(number phoneval.Number) (int, error)
}

// AreaCodeStatistics data struct with the statistics associated with an AreaCode
type AreaCodeStatistics struct {
	AreaCode int
	Quantity int
}

// PhoneStats phone statistics usecase
type PhoneStats struct {
	validator Validator
}

// NewPhoneStats factory function to instantiate a new PhoneStats usecase
func NewPhoneStats(validator phoneval.Validator) PhoneStats {
	return PhoneStats{
		validator: validator,
	}
}

// CalculateAreaCodeStatistics calculates AreaCode statistics and returns a slice of sorted AreaCodes and a map of statistics
func (p PhoneStats) CalculateAreaCodeStatistics(numbers []string) ([]int, map[int]int) {
	stats := map[int]int{}
	for _, n := range numbers {
		numb := phoneval.NewNumber(n)
		if isVal := p.validator.IsValid(numb); !isVal {
			log.Printf("invalid number found [%s]\n", n)
			continue
		}

		ac, err := p.validator.FindAreaCode(numb)
		if err != nil {
			log.Printf("error findind area code [%s][%s]\n", n, err)
			continue
		}

		stats[ac]++
	}

	return p.sortAreaCodes(stats), stats
}

func (p PhoneStats) sortAreaCodes(stats map[int]int) []int {
	keys := []int{}
	for ac := range stats {
		keys = append(keys, ac)
	}
	sort.Ints(keys)

	return keys
}
