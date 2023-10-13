package main

import (
	"errors"
	"fmt"
)

func checkParams(input string, span int) error {
	inputLen := len(input)
	if span < 0 {
		return errors.New("Span must be an non-negative.")
	} else if inputLen < span {
		return errors.New("Input less than span.")
	} else {
		for _, b := range []byte(input) {
			if b < '0' || b > '9' {
				return errors.New("Input contains non-digit.")
			}
		}
	}
	return nil
}
func LargestSeriesProduct(input string, span int) (int64, error) {
	err := checkParams(input, span)
	if err != nil {
		return -1, err
	}
	tmp := []byte(input)
	var max int64
	for i := 0; i <= len(tmp)-span; i++ {
		var product int64 = 1
		for j := i; j < i+span; j++ {
			n := int64(tmp[j] - '0')
			if n == 0 {
				product = 0
				i = j + 1
				break
			}
			product *= n
		}
		if product > max {
			max = product
		}
	}
	return max, nil
}
func main() {
	fmt.Println(LargestSeriesProduct("63915", 3))
}
