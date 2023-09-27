package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Order(sentence string) string {
	arr := strings.Fields(sentence)

	result := make([]string, len(arr))
	for j := 0; j < len(arr); j++ {
		for _, char := range arr[j] {

			if char <= 57 && char > 48 {
				v := string(char)

				if s, err := strconv.ParseInt(v, 10, 64); err == nil {
					result[s-1] = arr[j]
				}
			}
		}
	}
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(Order("4of Fo1r pe6ople g3ood th5e the2"))
	fmt.Println(Order("is2 Thi1s T4est 3a"))
}
