package roman

import "strings"

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for arabic > 0 {
		switch {
		case arabic >= 50:
			result.WriteString("L")
			arabic -= 50
		case arabic >= 40:
			result.WriteString("XL")
			arabic -= 40
		case arabic >= 10:
			result.WriteString("X")
			arabic -= 10
		case arabic >= 9:
			result.WriteString("IX")
			arabic -= 9
		case arabic >= 5:
			result.WriteString("V")
			arabic -= 5
		case arabic >= 4:
			result.WriteString("IV")
			arabic -= 4
		default:
			result.WriteString("I")
			arabic--
		}
	}

	return result.String()
}
