package main

import "fmt"

func Dominator(a []int) int {
	myMap := map[int]int{}
	for _, num := range a {
		myMap[num]++
	}
	var max int
	var dominator int
	for k, value := range myMap {
		if value > max {
			max = value
			dominator = k
		}
	}
	fmt.Println(myMap)
	if max > len(a)/2 {
		return dominator
	}
	return -1
}
func main() {
	arrTest := []int{3, 4, 3, 2, 3, 1, 3, 3}
	fmt.Println(Dominator(arrTest))
	arrTest2 := []int{1, 1, 1, 2, 2, 2}
	fmt.Println(Dominator(arrTest2))
}
