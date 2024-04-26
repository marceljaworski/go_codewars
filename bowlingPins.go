package main

import (
	"fmt"
	"strings"
)

func BowlingPins(slice []int) string {
	firstLine := make([]int, 0, 7)
	secondLine := make([]int, 0, 7)
	thirdLine := make([]int, 0, 7)
	fourthLine := make([]int, 0, 7)
	for i := 0; i <= 10; i++ {
		if i > 6 {
			if i == 10 {
				firstLine = append(firstLine, i)
			} else {
				firstLine = append(firstLine, i, 0)
			}
		}
		if i < 7 && i > 3 {
			if i == 6 {
				secondLine = append(secondLine, i)
			} else {
				secondLine = append(secondLine, i, 0)
			}
		}
		if i < 4 && i > 1 {
			if i == 3 {
				thirdLine = append(thirdLine, i)
			} else {
				thirdLine = append(thirdLine, i, 0)

			}
		}
		if i == 1 {
			fourthLine = append(fourthLine, i)
		}
	}
	secondLine = append([]int{0}, secondLine...)
	secondLine = append(secondLine, 0)

	thirdLine = append([]int{0, 0}, thirdLine...)
	thirdLine = append(thirdLine, []int{0, 0}...)

	fourthLine = append([]int{0, 0, 0}, fourthLine...)
	fourthLine = append(fourthLine, []int{0, 0, 0}...)

	var result strings.Builder
	result.WriteString(convertIntSliceToString(replace(firstLine, slice)) + "\n")
	result.WriteString(convertIntSliceToString(replace(secondLine, slice)) + "\n")
	result.WriteString(convertIntSliceToString(replace(thirdLine, slice)) + "\n")
	result.WriteString(convertIntSliceToString(replace(fourthLine, slice)))

	return result.String()
}

func replace(pins []int, numbers []int) []int {
	for _, number := range numbers {
		for i, pin := range pins {
			if number == pin {
				pins[i] = 0
			}
		}
	}
	return pins
}
func convertIntSliceToString(line []int) string {
	var result = make([]string, 0, 7)
	for _, intValue := range line {
		if intValue == 0 {
			result = append(result, " ")
		} else {
			result = append(result, "I")
		}
	}
	return strings.Join(result, "")
}
func main() {
	fmt.Println(BowlingPins([]int{3, 5, 9}))
}
