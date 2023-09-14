package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	// newStr := strings.ToLower(str)
	// var result string
	// for _, v := range newStr {
	// 	result = string(v) + result
	// }
	// fmt.Println(newStr, result)
	// return newStr == result
	str = strings.ToLower(str)
	n := len(str)
	for i := 0; i < n; i++ {
		n -= 1
		if str[i] != str[n] {
			return false
		}
	}
	return true
}
func main() {

	fmt.Println(IsPalindrome("Neveroddoreven"))
}
