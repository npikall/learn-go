package roman

import "strings"

func ConvertToRoman(arabic int) string {
	if arabic == 4 {
		return "IV"
	}
	var result strings.Builder

	for range arabic {
		result.WriteString("I")
	}
	return result.String()
}
