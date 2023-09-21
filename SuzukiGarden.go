package main

import (
	"fmt"
	"strings"
)

func RakeGarden(garden string) string {
	str := strings.Fields(garden)
	for i, field := range str {

		if field != "gravel" || field != "rock" {
			str[i] = "gravel"
		}
	}

	return strings.Join(str, " ")
}

func main() {
	fmt.Println(RakeGarden("slug spider rock gravel gravel gravel gravel gravel gravel gravel"))
}
