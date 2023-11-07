package main

import (
	"fmt"
)

// Use the preloaded Tuple struct as return type
// type Tuple struct {
// 	Char  rune
// 	Count int
// }

// func OrderedCount(text string) []Tuple {
// 	res := make([]Tuple, 0, len(text))
// 	myMap := make(map[rune]struct{})

//		for _, char := range text {
//			if _, ok := myMap[char]; !ok {
//				res = append(res, Tuple{Char: char, Count: strings.Count(text, string(char))})
//				myMap[char] = struct{}{}
//			}
//		}
//		return res
//	}
func OrderedCount(text string) []Tuple {
	dict := make(map[rune]int)
	order := []rune{}
	for _, rune := range text {
		if v, ok := dict[rune]; ok {
			dict[rune] = v + 1
		} else {
			dict[rune] = 1
			order = append(order, rune)
		}
	}
	out := make([]Tuple, len(order))

	for i, k := range order {
		out[i] = Tuple{k, dict[k]}
	}
	return out
}
func main() {
	fmt.Println(OrderedCount("abracadabra"))
}
