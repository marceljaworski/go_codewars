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
				fmt.Println(v, arr[j])
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

// You can edit this code!
// Click here and start typing.
// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	input := "4of Fo0r pe10ople g3ood th5e the2"
// 	arr := strings.Fields(input)

// 	sort.Sort(ByNumber(arr))

// 	fmt.Println(strings.Join(arr, " "))
// }

// type ByNumber []string

// func (array ByNumber) Len() int {
// 	return len(array)
// }

// func (array ByNumber) Less(i, j int) bool {
// 	return findNumberInString(array[i]) < findNumberInString(array[j])
// }

// func findNumberInString(val string) int {
// 	var rslt string
// 	for _, char := range val {
// 		if rslt != "" && (char < 48 || char > 57) {
// 			break
// 		}
// 		if char >= 48 && char <= 57 {
// 			rslt += string(char)
// 		}
// 	}
// 	out, _ := strconv.Atoi(rslt)
// 	return out
// }

// func (array ByNumber) Swap(i, j int) {
// 	array[i], array[j] = array[j], array[i]
// }
