package main

import "fmt"

func SequenceSum(start, end, step int) int {
	count := 0
	for i := start; i <= end; i += step {
		count += i
	}
	// if start > end {
	// 	return 0
	// }
	// slice := []int{start}
	// count := start
	// for count+step <= end {
	// 	count = count + step

	// 	slice = append(slice, count)
	// }
	// solution := 0
	// for _, num := range slice {
	// 	solution = solution + num
	// }
	// return solution
	return count
}

func main() {
	fmt.Println(SequenceSum(2, 6, 2))
	fmt.Println(SequenceSum(1, 5, 3))
}
