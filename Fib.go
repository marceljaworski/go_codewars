package main

import "fmt"

func Fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		c := b
		b = a + b
		a = c
	}
	return a
}
func main() {
	fmt.Println(Fib(4))
}
