package main

import (
	"fmt"
	"strings"
)

func ToJadenCase(str string) string {
	words := strings.Fields(str)
	for i := 0; i < len(words); i++ {
		chars := strings.Split(words[i], "")
		chars[0] = strings.ToUpper(chars[0])
		words[i] = strings.Join(chars, "")
		fmt.Println(words[i])
	}

	return strings.Join(words, " ")
}
func main() {
	fmt.Println(ToJadenCase("should work for sample test cases"))
	fmt.Println(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own."))

}
