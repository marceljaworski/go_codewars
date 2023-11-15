package main

import (
	"fmt"
	"regexp"
)

func ReverseLetters(s string) string {
	re := regexp.MustCompile(`[^a-z]`)
	replaced := re.ReplaceAllString(s, "")
	solution := ""
	for i := len(replaced) - 1; i >= 0; i-- {
		solution = solution + string(replaced[i])
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
