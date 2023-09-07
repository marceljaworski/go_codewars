package main

import "fmt"

func CalculateYears(years int) (result [3]int) {
	catYears, dogYears := 15, 15

	for i := 1; i < years; i++ {
		if i == 1 {
			catYears += 9
			dogYears += 9
		} else {
			catYears += 4
			dogYears += 5
		}
	}
	return [3]int{years, catYears, dogYears}
}

func main() {
	fmt.Println(CalculateYears(1))
	fmt.Println(CalculateYears(2))
	fmt.Println(CalculateYears(3))
	fmt.Println(CalculateYears(4))
}
