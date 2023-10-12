package main

import (
	"fmt"
	"regexp"
	"strings"
)

func CheckPairs(nbBear int, bearList string) (string, bool) {
	re := regexp.MustCompile("B8|8B").FindAllString(bearList, -1)
	fmt.Println(re)
	response := strings.Join(re, "")
	return response, len(response) >= nbBear
	// male := ""
	// female := ""
	// var response string
	// for _, bear := range bearList {
	// 	if bear == 66 {
	// 		male += string(bear)
	// 	}
	// 	if bear == 56 {
	// 		female += string(bear)
	// 	}
	// }
	// if len(male) != len(female) {
	// 	if len(male) > len(female) {
	// 		rest := len(male) - len(female)
	// 		male = male[:len(male)-rest]
	// 	} else {
	// 		rest := len(female) - len(male)
	// 		female = female[:len(female)-rest]
	// 	}
	// }
	// response = male + female

	// return response, len(male) >= nbBear

}
func main() {
	fmt.Println(CheckPairs(1, "j8BmB88B88gkBBlf8hg8888lbe88"))
}
