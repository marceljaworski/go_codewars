package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Revrot(s string, n int) string {
	if len(s) == 0 || n == 0 || n > len(s) {
		return ""
	}
	chunks := []string{}
	for i := 0; i+n <= len(s); i += n {
		chunks = append(chunks, reverseOrRotate(s[i:i+n]))
	}
	return strings.Join(chunks, "")
}
func reverseOrRotate(s string) string {
	var count float64
	// sum of the cubes of its digits
	for _, char := range s {
		num, _ := strconv.Atoi(string(char))
		count += math.Pow(float64(num), 3)
	}
	solution := ""
	if math.Mod(count, 2) == 0 {
		//reverse string
		for i := len(s) - 1; i >= 0; i-- {
			solution += string(s[i])
		}
	} else {
		// rotate string
		solution += s[1:] + s[:1]
	}
	return solution
}
func main() {
	// fmt.Println(Revrot("123456987654", 6))
	fmt.Println(Revrot("733049910872815764", 5)) //33047 91089 28157
	fmt.Println(Revrot("773256782953519667796", 7))
}
