package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Valid(id string) bool {

	idStr := strings.ReplaceAll(id, " ", "")
	count := 0
	if len(idStr) < 2 {
		return false
	}
	//Reverse the string
	runes := []rune(idStr)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	slice := make([]int, 0, len(runes))
	for i := 0; i < len(runes); i++ {
		num, err := strconv.Atoi(string(runes[i]))
		if err != nil {
			return false
		}
		slice = append(slice, num)
	}

	for i, number := range slice {
		if i%2 != 0 {
			if number*2 > 9 {
				count += number*2 - 9
			} else {
				count += number * 2
			}
		} else {
			count += number
		}
	}
	if count%10 == 0 {
		return true
	}
	return false
}
func main() {
	fmt.Println(Valid("055b 444 285"))
}
