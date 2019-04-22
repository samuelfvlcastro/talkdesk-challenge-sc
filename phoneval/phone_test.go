package phoneval

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestNewNumber(t *testing.T) {
	RegisterTestingT(t)

	testCases := []struct {
		desc     string
		number   string
		expected Number
	}{
		{
			"number without formatting gets cleaned",
			"916454099",
			Number{
				Original: "916454099",
				Clean:    "916454099",
			},
		},
		{
			"number with initial +",
			"+916454099",
			Number{
				Original: "+916454099",
				Clean:    "916454099",
			},
		},
		{
			"number with initial double zeros",
			"00916454099",
			Number{
				Original: "00916454099",
				Clean:    "916454099",
			},
		},
		{
			"number with initial triple zeros",
			"00016454099",
			Number{
				Original: "00016454099",
				Clean:    "016454099",
			},
		},
		{
			"number with initial zeros and spaces between numbers gets cleaned",
			"0091 645 40 99",
			Number{
				Original: "0091 645 40 99",
				Clean:    "916454099",
			},
		},
		{
			"number with initial + and spaces between numbers gets cleaned",
			"+91 645 40 99",
			Number{
				Original: "+91 645 40 99",
				Clean:    "916454099",
			},
		},
		{
			"invalid number with plus and two zeros still gets cleaned",
			"+00916454099",
			Number{
				Original: "+00916454099",
				Clean:    "916454099",
			},
		},
		{
			"invalid number with characters different than '+' doesn't get cleaned",
			"+91-645-40-99t",
			Number{
				Original: "+91-645-40-99t",
				Clean:    "91-645-40-99t",
			},
		},
	}

	for _, tC := range testCases {
		numb := NewNumber(tC.number)

		Expect(numb).To(Equal(tC.expected), tC.desc)
	}
}
