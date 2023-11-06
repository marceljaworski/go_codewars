package main

import "fmt"

func MxDifLg(a1 []string, a2 []string) int {
	maxA1 := 0
	maxA2 := 0

	for _, str1 := range a1 {
		if len(str1) > maxA1 {
			maxA1 = len(str1)
		}
	}
	for _, str2 := range a2 {
		if len(str2) > maxA2 {
			maxA2 = len(str2)
		}
	}
	return maxA1 - maxA2
}
func main() {
	var s1 = []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	var s2 = []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	// var s1 = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	// var s2 = []string{"bbbaaayddqbbrrrv"}
	fmt.Println(MxDifLg(s1, s2))
}
