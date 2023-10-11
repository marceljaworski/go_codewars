package main

import (
	"errors"
	"fmt"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) || span < 0 {
		return 0, errors.New("Error")
	}
	slice := make([]string, 0, span)
	for i := 0; i < span; i++ {
		slice = append(slice, digits[i:i+span])
	}
	var response int64
	for j := 0; j < len(slice); j++ {
		count := 0
		for i := 0; i < len(slice[j]); i++ {
			v, _ := strconv.Atoi(string(slice[j][i]))
			if i == 0 {
				count += v
			} else {
				count = count * v
			}
		}
		if response < int64(count) {
			response = int64(count)
		}
	}
	return response, nil
}
func main() {
	fmt.Println(LargestSeriesProduct("63915", 3))
}
