package main

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	slice := strings.Split(strings.ReplaceAll(isbn, "-", ""), "")
	count := 0
	multiple := 10
	if len(slice) != 10 {
		return false
	}
	if slice[len(slice)-1] == "X" || slice[len(slice)-1] == "B" {
		slice[len(slice)-1] = "10"
	}
	for _, number := range slice {
		num, err := strconv.Atoi(number)
		if err != nil {
			return false
		}
		count += num * multiple
		multiple--
	}
	if count%11 == 0 {
		return true
	}
	return false
}
func main() {
	IsValidISBN("98245726788")
}
