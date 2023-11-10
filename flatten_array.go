package main

import "fmt"

func Flatten(nested interface{}) []interface{} {
	solution := []interface{}{}
	switch s := nested.(type) {
	case []interface{}:
		for _, val := range s {
			solution = append(solution, Flatten(val)...)
		}
	case interface{}:
		solution = append(solution, s)
	}
	return solution
}

// func sliceFunction(s []interface{}) (solution []interface{}) {
// 	if len(s) == 0 {
// 		return s
// 	}
// 	for _, value := range s {
// 		switch value2 := value.(type) {
// 		case []interface{}:
// 			sliceFunction(value2)
// 		case nil:
// 			continue
// 		case int:
// 			solution = append(solution, s)
// 		}
// 	}
// 	return
// }
// func mapsFunction(s map[int]interface{}) (solution []interface{}) {
// 	if len(s) == 0 {
// 		return
// 	}
// 	for _, value := range s {
// 		switch value2 := value.(type) {
// 		case []interface{}:
// 			sliceFunction(value2)
// 		case nil:
// 			continue
// 		case int:
// 			solution = append(solution, s)
// 		}
// 	}
// 	return
// }

//	func appendFunction(s int) (solution []interface{}) {
//		solution = append(solution, s)
//		return solution
//	}
func main() {
	fmt.Println(Flatten([]int{0, 1, 2}))
}
