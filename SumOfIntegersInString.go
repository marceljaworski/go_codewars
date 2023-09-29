package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SumOfIntegersInString(str string) int {
	slice := []string{}
	sum := 0
	for i := 0; i < len(str); i++ {
		if strings.ContainsAny(string(str[i]), "1234567890") {
			slice = append(slice, string(str[i]))
		} else {
			if s, err := strconv.Atoi(strings.Join(slice, "")); err == nil {
				sum += s
			}
			slice = slice[:0]
		}
	}
	if len(slice) > 0 {
		if s, err := strconv.Atoi(strings.Join(slice, "")); err == nil {
			sum += s
		}
	}
	return sum
}

func main() {
	fmt.Println(SumOfIntegersInString("h3ll0w0rld"))
	fmt.Println(SumOfIntegersInString("12.4"))
	fmt.Println(SumOfIntegersInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"))
}
