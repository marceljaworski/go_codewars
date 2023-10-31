package main

import (
	"fmt"
	"regexp"
)

func ReverseLetters(s string) string {
	re := regexp.MustCompile(`[^a-z]`)
	replaced := re.ReplaceAllString(s, "")
	count := len(replaced) - 1
	solution := ""
	for i := 0; i < len(replaced); i++ {
		solution = solution + string(replaced[count])
		count--
	}
	return solution
}

// func ReverseLetters(s string) string {
// 	arr := []rune{}
// 	for _, v := range s {
// 		if unicode.IsLetter(v) {
// 			arr = append([]rune{v}, arr...)
// 		}
// 	}
// 	return string(arr)

// }
func main() {
	fmt.Println(ReverseLetters("ultr53o?n"))
}
