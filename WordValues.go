package main

import "fmt"

func NameValue(my_list []string) []int {
	count_list := make([]int, 0, len(my_list))
	for i, word := range my_list {
		count := 0
		for _, char := range word {
			if char > 96 && char < 123 {
				count = count + (int(char) - 96)
			}
		}
		count_list = append(count_list, count*(i+1))
	}
	return count_list
}
func main() {
	fmt.Println(NameValue([]string{"abc", "abc", "abc", "abc"}))
}
