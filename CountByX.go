package main

import (
	"fmt"
)

func CountBy(x, n int) []int {

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = x * (i + 1)
		// // if i == n {
		// // 	break
		// // }
		// if x == 1 {
		// 	arr = append(arr, i)
		// } else {
		// 	arr = append(arr, i*x)
		// }

	}
	return arr
}
func main() {
	fmt.Println(CountBy(1, 10))
}
