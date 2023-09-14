package main

import (
	"fmt"
	"strconv"
)

func SumMix(arr []any) int {
	sum := 0
	for _, val := range arr {
		switch val := val.(type) {
		case int:
			sum += val

		case string:
			number, _ := strconv.Atoi(val)
			sum += number
		}

	}
	return sum
}

func main() {
	arr := []interface{}{"9", 5}
	fmt.Println(SumMix(arr))
}

//The most common numeric conversions are Atoi (string to int) and Itoa (int to string).

//ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values
