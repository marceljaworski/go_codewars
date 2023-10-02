package main

import (
	"fmt"
)

func FindUniq(arr []float32) float32 {

	var result float32 = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] != result {
			result = arr[i]
			return result
		} else {
			continue
		}
	}
	return result
}
func main() {
	arr := []float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}
	fmt.Println(FindUniq(arr))
}
