package main

import "fmt"

func BubblesortOnce(numbers []int) []int {
	list := make([]int, len(numbers))
	copy(list, numbers)
	for i := 1; i < len(list); i++ {
		if list[i-1] > list[i] {
			temp := list[i]
			list[i] = list[i-1]
			list[i-1] = temp
		}
	}
	return list
}
func main() {
	testArr := []int{7, 9, 5, 3, 1, 2, 4, 6, 8}
	fmt.Println(BubblesortOnce(testArr))
}
