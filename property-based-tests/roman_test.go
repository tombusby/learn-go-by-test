package roman

import (
	"fmt"
	"testing"
	"testing/quick"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var testCases = []RomanNumeral{
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
	{90, "XC"},
	{95, "XCV"},
	{98, "XCVIII"},
	{99, "XCIX"},
	{100, "C"},
	{150, "CL"},
	{399, "CCCXCIX"},
	{400, "CD"},
	{500, "D"},
	{700, "DCC"},
	{800, "DCCC"},
	{895, "DCCCXCV"},
	{900, "CM"},
	{1000, "M"},
	{1001, "MI"},
	{1051, "MLI"},
	{5000, "MMMMM"},
	{5905, "MMMMMCMV"},
}

func TestRomanNumerals(t *testing.T) {

	for _, test := range testCases {
		t.Run(fmt.Sprintf("%d gets converted to %s", test.Value, test.Symbol),
			func(t *testing.T) {
				got := ConvertToRoman(test.Value)
				if got != test.Symbol {
					t.Errorf("got %q, want %q", got, test.Symbol)
				}
			})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Symbol, test.Value), func(t *testing.T) {
			got := ConvertToArabic(test.Symbol)
			if got != test.Value {
				t.Errorf("got %d, want %d", got, test.Value)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
