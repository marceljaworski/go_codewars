package main

import (
	"fmt"
	"strings"
)

func SolveAbc(str string) int {
	var abc string = "aabcdefghijklmnopqrstuvwxyz"

	countConsonants := 0
	subStrings := make([]int, 0, len(str))
	for i := 0; i < len(str); i++ {

		if strings.Contains("aeiou", string(str[i])) {
			subStrings = append(subStrings, countConsonants)
			countConsonants = 0
			continue
		} else {
			countConsonants = countConsonants + strings.Index(abc, string(str[i]))
		}
	}
	subStrings = append(subStrings, countConsonants)
	max := subStrings[0]
	for _, number := range subStrings {
		if number > max {
			max = number
		}
	}
	return max

}

func main() {
	fmt.Println(SolveAbc("strenght")) // 57
	fmt.Println(SolveAbc("zodiacs"))  // 26
}
