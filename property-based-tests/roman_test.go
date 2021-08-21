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
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{13, "XIII"},
		{14, "XIV"},
		{15, "XV"},
		{18, "XVIII"},
		{19, "XIX"},
		{20, "XX"},
		{30, "XXX"},
		{39, "XXXIX"},
		{40, "XL"},
		{41, "XLI"},
		{45, "XLV"},
		{46, "XLVI"},
		{48, "XLVIII"},
		{49, "XLIX"},
		{50, "L"},
		{60, "LX"},
		{89, "LXXXIX"},
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
