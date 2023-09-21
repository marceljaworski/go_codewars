package main

import (
	"fmt"
	"sort"
)

func Gimme(array [3]int) int {
	arrayCopy := array

	arrSlice := arrayCopy[:]
	sort.Ints(arrSlice)
	val := arrSlice[1]

	for i, num := range array {

		if val == num {
			return i
		}
	}
	return 0
}

func main() {
	var arr = [3]int{2, 3, 1}
	var arr1 = [3]int{5, 10, 14}
	var arr2 = [3]int{1, 3, 4}
	var arr3 = [3]int{15, 10, 14}
	fmt.Println(Gimme(arr))
	fmt.Println(Gimme(arr1))
	fmt.Println(Gimme(arr2))
	fmt.Println(Gimme(arr3))
}
