package main

//Your task in this kata is to implement a function that calculates the sum of the integers inside a string.
// For example, in the string "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog", the sum of the integers is 3635.
import (
	"fmt"
	"regexp"
	"strconv"
	// "strings"
)

func SumOfIntegersInString(strng string) int {
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re, re.FindAllString(strng, -1))

	sum := 0
	for _, i := range re.FindAllString(strng, -1) {
		number, _ := strconv.Atoi(i)
		sum += number
	}
	return sum
}

// func SumOfIntegersInString(str string) int {
// 	slice := []string{}
// 	sum := 0
// 	for i := 0; i < len(str); i++ {
// 		// if strings.ContainsAny(string(str[i]), "1234567890") {
// 		if strings.ContainsAny(string(str[i]), "1234567890") {
// 			slice = append(slice, string(str[i]))
// 		} else {
// 			if s, err := strconv.Atoi(strings.Join(slice, "")); err == nil {
// 				sum += s
// 			}
// 			slice = slice[:0]
// 		}
// 	}
// 	if len(slice) > 0 {
// 		if s, err := strconv.Atoi(strings.Join(slice, "")); err == nil {
// 			sum += s
// 		}
// 	}
// 	return sum
// }

func main() {
	fmt.Println(SumOfIntegersInString("h3ll0w0rld"))
	// fmt.Println(SumOfIntegersInString("12.4"))
	// fmt.Println(SumOfIntegersInString("The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"))
}
