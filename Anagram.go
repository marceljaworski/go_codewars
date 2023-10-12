package main

import (
	"fmt"
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	str := strings.ToLower(subject)
	response := make([]string, 0, len(candidates))
	subjectMap := map[rune]int{}
	for _, char := range str {
		subjectMap[char]++
	}
	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)
		candidateMap := map[rune]int{}
		for _, char := range candidateLower {
			candidateMap[char]++
		}
		eq := reflect.DeepEqual(subjectMap, candidateMap)
		if str != candidateLower && eq {
			response = append(response, candidate)
		}
	}
	return response
}
func main() {
	testSlice := []string{"hello", "world", "zombies", "pants"}
	// test2 := []string{"stream", "pigeon", "maters"}
	fmt.Println(Detect("diaper", testSlice))
	// fmt.Println(Detect("master", test2))
}
