package roman

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		Arabic int
		Want   string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %s", test.Arabic, test.Want),
			func(t *testing.T) {
				got := ConvertToRoman(test.Arabic)
				if got != test.Want {
					t.Errorf("got %q, want %q", got, test.Want)
				}
			})
	}
}
