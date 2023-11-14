package main

import (
	"fmt"
	"math"
)

func PowersOfTwo(n int) []uint64 {
	slice := []uint64{}
	for i := 0; i <= n; i++ {
		c := math.Pow(float64(2), float64(n)
		slice = append(slice, uint64(c))
	}
	return slice
}
func main() {
	fmt.Println(PowersOfTwo(0))
}
