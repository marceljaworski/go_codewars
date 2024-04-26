package main

import (
	"fmt"
	"strings"
)

func SortVowels(s string) string {
	var buffer strings.Builder
	if s == "" {
		return ""
	}
	for _, char := range s {
		switch char {
		case 97:
			buffer.WriteString(printf(char))
		case 65:
			buffer.WriteString(printf(char))
		case 69:
			buffer.WriteString(printf(char))
		case 101:
			buffer.WriteString(printf(char))
		case 73:
			buffer.WriteString(printf(char))
		case 105:
			buffer.WriteString(printf(char))
		case 79:
			buffer.WriteString(printf(char))
		case 111:
			buffer.WriteString(printf(char))
		case 85:
			buffer.WriteString(printf(char))
		case 117:
			buffer.WriteString(printf(char))
		default:
			buffer.WriteString(fmt.Sprintf("%s|\n", string(char)))

		}
	}
	return strings.TrimSuffix(buffer.String(), "\n")

}
func printf(char rune) string {
	return fmt.Sprintf("|%s\n", string(char))
}
func main() {
	fmt.Println(SortVowels("Code Wars"))
}
