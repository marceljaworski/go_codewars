package main

import (
	"errors"
	"fmt"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("input number is equal or less than zero")
	}
	primes := []int{}
	pn := 0
	i := 0
	for ; pn != n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
			pn++
		}
	}
	fmt.Println(primes)
	return primes[n-1], nil
}
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}
func main() {
	fmt.Println(Nth(6))
}
