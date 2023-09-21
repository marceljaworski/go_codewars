package main

import "fmt"

func ExpressionMatter(a int, b int, c int) int {

	compare := [5]int{
		a * (b + c),
		a * b * c,
		a + b*c,
		(a + b) * c,
		a + b + c,
	}
	var max int

	for _, v := range compare {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(ExpressionMatter(2, 1, 2)) //6
	fmt.Println(ExpressionMatter(1, 1, 1)) //3
}
