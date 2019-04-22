package phoneval

import (
	"testing"

	. "github.com/onsi/gomega"
	"gitlab.com/samuelfvlcastro/talkdesk-challenge-sc/x/rtrie"
)

func TestNewValidator(t *testing.T) {
	RegisterTestingT(t)

	val := NewValidator([]string{"345"})

	expTrie := rtrie.New(48, 57)
	expTrie.Insert("345")
	expVal := Validator{
		expTrie,
	}

	Expect(val).To(Equal(expVal), "expect to return a new validator instance with the correct trie")
}

func TestIsValid(t *testing.T) {
	RegisterTestingT(t)

	val := NewValidator([]string{})

	testCases := []struct {
		desc     string
		number   Number
		expected OmegaMatcher
	}{
		{
			"valid number without prefixes",
			Number{
				Clean:    "916454099",
				Original: "916454099",
			},
			BeTrue(),
		},
		{
			"valid small number with prefixes",
			Number{
				Clean:    "112",
				Original: "00112",
			},
			BeTrue(),
		},
		{
			"valid number with +",
			Number{
				Clean:    "916454099",
				Original: "+916454099",
			},
			BeTrue(),
		},
		{
			"valid number with double zeros",
			Number{
				Clean:    "916454099",
				Original: "00916454099",
			},
			BeTrue(),
		},
		{
			"valid number with double zeros and spaces between",
			Number{
				Clean:    "916454099",
				Original: "0091 645 40 99",
			},
			BeTrue(),
		},
		{
			"valid number with + and spaces between",
			Number{
				Clean:    "916454099",
				Original: "+91 645 40 99",
			},
			BeTrue(),
		},
		{
			"invalid number with + and extra zeros",
			Number{
				Clean:    "916454099",
				Original: "+00916454099",
			},
			BeFalse(),
		},
		{
			"invalid number with extra numbers after",
			Number{
				Clean:    "12",
				Original: "0012",
			},
			BeFalse(),
		},
		{
			"invalid number with extra numbers before",
			Number{
				Clean:    "35112",
				Original: "35112",
			},
			BeFalse(),
		},
		{
			"invalid number with both the + and 00 prefixes",
			Number{
				Clean:    "112",
				Original: "+00112",
			},
			BeFalse(),
		},
		{
			"invalid number with space after the +",
			Number{
				Clean:    "916454099",
				Original: "+ 916454099",
			},
			BeFalse(),
		},
		{
			"invalid number invalid characters between the number",
			Number{
				Clean:    "9164t54099",
				Original: "9164t54099",
			},
			BeFalse(),
		},
	}

	for _, tC := range testCases {
		Expect(val.IsValid(tC.number)).To(tC.expected, tC.desc)
	}
}

func TestFindAreaCode(t *testing.T) {
	RegisterTestingT(t)

	val := NewValidator([]string{"91", "912", "9123", "8"})

	testCases := []struct {
		desc     string
		number   Number
		err      OmegaMatcher
		expected int
	}{
		{
			"existing area code 91",
			Number{
				Clean:    "910000000",
				Original: "910000000",
			},
			Not(HaveOccurred()),
			91,
		},
		{
			"existing area code 912",
			Number{
				Clean:    "912000000",
				Original: "912000000",
			},
			Not(HaveOccurred()),
			912,
		},
		{
			"existing area code 9123",
			Number{
				Clean:    "912300000",
				Original: "912300000",
			},
			Not(HaveOccurred()),
			9123,
		},
		{
			"existing area code 8",
			Number{
				Clean:    "812300000",
				Original: "812300000",
			},
			Not(HaveOccurred()),
			8,
		},
		{
			"non-existing area code",
			Number{
				Clean:    "923400000",
				Original: "923400000",
			},
			Not(HaveOccurred()),
			0,
		}, {
			"invalid rune throws error",
			Number{
				Clean:    "9+23000000",
				Original: "9+23000000",
			},
			HaveOccurred(),
			0,
		},
	}

	for _, tC := range testCases {

		ac, err := val.FindAreaCode(tC.number)
		Expect(err).To(tC.err)
		Expect(ac).To(Equal(tC.expected), tC.desc)
	}

}
