package main

import (
	"fmt"
	"strings"
	"unicode"
)

// values" hand " "hund" " hund und " "und"
// delim"  und  "
// "hand  und  hund  und  hund"

func trimPreSuffix(delimiter, value string) string {
	value = strings.TrimSpace(value)
	// fmt.Printf(`"%s" - %v%s`, delimiter, value[len(value)-1], "\n")
	if delimiter != "" {
		delimRunes := []rune(delimiter)
		//Trim Suffix
		if unicode.IsSpace(delimRunes[0]) {
			trimmed := strings.TrimSuffix(value, strings.TrimSpace(delimiter))
			if strings.TrimSpace(trimmed) == "" {
				return ""
			}
			if trimmed != value {
				if runes := []rune(trimmed); unicode.IsSpace(runes[len(runes)-1]) {
					value = strings.TrimSpace(trimmed)
				}
			}

		} else {
			value = strings.TrimSpace(strings.TrimSuffix(value, strings.TrimSpace(delimiter)))
		}
		// Trim Prefix
		if unicode.IsSpace(delimRunes[len(delimRunes)-1]) {
			trimmed := strings.TrimPrefix(value, strings.TrimSpace(delimiter))
			if strings.TrimSpace(trimmed) == "" {
				return ""
			}
			if trimmed != value {
				if runes := []rune(trimmed); unicode.IsSpace(runes[0]) {
					value = strings.TrimSpace(trimmed)
				}
			}

		} else {
			value = strings.TrimSpace(strings.TrimPrefix(value, strings.TrimSpace(delimiter)))
		}
	}
	return value
}

func debug(delim string, values ...string) {
	for _, value := range values {
		fmt.Printf("Delimiter: \"%s\" - Value: \"%s\" - Result: \"%s\"\n", delim, value, trimPreSuffix(delim, value))
	}
}

func main() {
	// debug(" ", " hallo ")
	debug(" , ", " hallo ", " , hallo", ", hallo", " ,hallo")
	debug("/", " a ", "/b", "/ c /", " hallo")
	debug(" und ", " und ", "hund und", "hand", " und hund")
}
