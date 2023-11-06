package main

import "fmt"

func ArrayDiff(a, b []int) []int {
	bKeys := make(map[int]bool)
	slice := make([]int, 0, len(a))
	for _, bItem := range b {
		bKeys[bItem] = true
	}
	for _, aItem := range a {
		if bKeys[aItem] == false {
			slice = append(slice, aItem)
		}
	}
	return slice
}
func main() {
	a := []int{1, 2, 2, 3}
	b := []int{2}
	fmt.Println(ArrayDiff(a, b))
}
