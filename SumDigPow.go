package main

import (
	"fmt"
	"strconv"
)

func SumDigPow(a, b uint64) []uint64 {
	slice := []uint64{}
	for i := a; i <= b; i++ {

		s := strconv.FormatUint(i, 10)
		var count uint64 = 0
		for j, digit := range s {
			num := uint64(digit - '0')
			pow := num
			for k := 0; k < j; k++ {
				pow *= num
			}
			count += pow
		}
		if count == i {
			slice = append(slice, i)
		}
	}
	return slice
}
func main() {
	fmt.Println(SumDigPow(1, 100))
	// fmt.Println(SumDigPow(12157692622039623120, 12157692622039625680))
	// fmt.Println(SumDigPow(12157692622039623539, 12157692622039623539))
	// 12157665460000000000
}
