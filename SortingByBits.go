package main

import (
	"fmt"
	"math/bits"
	"sort"
)

// func getOne(n int) (count int) {
// 	b := strconv.FormatInt(int64(n), 2)
// 	for _, c := range b {
// 		if c == '1' {
// 			count++
// 		}
// 	}
// 	return
// }

func SortByBit(arr []int) {
	// sort the input array in place
	sort.Slice(arr, func(i, j int) bool {
		bi, bj := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		if bi == bj {
			return arr[i] < arr[j]
		}
		return bi < bj
	})

}
func main() {
	test := []int{3, 8, 3, 6, 5, 7, 9, 1}
	fmt.Println(SortByBit(test)) // --> [1, 8, 3, 3, 5, 6, 9, 7]
}
