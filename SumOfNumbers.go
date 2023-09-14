package main

import "fmt"

func GetSum(a, b int) (sum int) {
	if a == b {
		return b
	}

	if a > b {
		for i := b; i <= a; i++ {

			sum += i
		}
	}
	if a < b {
		for i := a; i <= b; i++ {

			sum += i
		}
	}
	return
}

func main() {
	fmt.Println(GetSum(0, 1))
	fmt.Println(GetSum(1, 2))
	fmt.Println(GetSum(5, -1))
	fmt.Println(GetSum(505, 4))
}

// https://www.codewars.com/kata/55f2b110f61eb01779000053/go
