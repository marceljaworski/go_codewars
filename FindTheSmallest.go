package main

import (
	"fmt"
	"strconv"
	"strings"
)

func smallest(n int64) []int64 {
	arr := make([]int64, 3, 3)
	s := strconv.Itoa(int(n))
	slice := strings.Split(s, "")
	var min rune
	min = 57
	var index int
	for i, run := range s {
		if min > run {
			min = run
			index = i
		}
	}
	slice = append(slice[:index], slice[index+1:]...)
	slice = append([]string{string(min)}, slice...)
	i, err := strconv.ParseInt(strings.Join(slice, ""), 10, 64)
	if err != nil {
		panic(err)
	}
	arr[0] = int64(i)
	arr[1] = int64(index)
	arr[2] = int64(0)
	fmt.Println(slice)
	if arr[1] > arr[2] {
		arr[1] = int64(0)
		arr[2] = int64(index)
	}
	return arr
}
func main() {
	fmt.Println(smallest(261235))
	fmt.Println(smallest(209917))
}
