package main

import (
	"fmt"
	"strings"
	"unicode"
)

func BandNameGenerator(word string) string {
	var solution = make([]string, 0, 2)
	runes := []rune(word)
	var name string = ""
	if string(runes[len(runes)-1:]) != string(runes[0]) {

		solution = append(solution, "The")

		if strings.Contains(word, "-") {

			splited := strings.Split(word, "-")
			var solution1 = make([]string, 0, len(splited))
			for _, word1 := range splited {
				runes1 := []rune(word1)
				runes1[0] = unicode.ToUpper(runes1[0])
				solution1 = append(solution1, string(runes1))
			}
			name1 := strings.Join(solution1, "-")
			solution = append(solution, name1)
		} else {
			runes[0] = unicode.ToUpper(runes[0])
			solution = append(solution, string(runes))

		}

		name = strings.Join(solution, " ")
	} else {

		runes[0] = unicode.ToUpper(runes[0])
		solution = append(solution, string(runes))
		sufix := runes[1:]
		solution = append(solution, string(sufix))
		name = strings.Join(solution, "")
	}
	return name
}

func main() {
	fmt.Println(BandNameGenerator("riders"))
	fmt.Println(BandNameGenerator("riders-are-cool"))

	fmt.Println(BandNameGenerator("alaska"))
}

//https://www.codewars.com/kata/59727ff285281a44e3000011/train/go
