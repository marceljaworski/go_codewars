package main

import "fmt"

func DirReduc(arr []string) []string {

	tmpArr := []string{}

	for i := len(arr) - 1; i > 0; i-- {
		if (arr[i] == "NORTH" && arr[i-1] == "SOUTH") ||
			(arr[i] == "SOUTH" && arr[i-1] == "NORTH") ||
			(arr[i] == "EAST" && arr[i-1] == "WEST") ||
			(arr[i] == "WEST" && arr[i-1] == "EAST") {
			tmpArr = append(tmpArr, arr[:i-1]...)
			tmpArr = append(tmpArr, arr[i+1:]...)
			return DirReduc(tmpArr)
		}
	}

	return arr
}

// func checkDirections(arrCopy []string) []string {
// 	for i := 1; i < len(arrCopy); i++ {
// 		if arrCopy[i-1] == "NORTH" && arrCopy[i] == "SOUTH" {
// 			arrCopy = append(arrCopy[:i-1], arrCopy[i+1:]...)
// 		}
// 		if arrCopy[i-1] == "SOUTH" && arrCopy[i] == "NORTH" {
// 			arrCopy = append(arrCopy[:i-1], arrCopy[i+1:]...)
// 		}
// 		if arrCopy[i-1] == "EAST" && arrCopy[i] == "WEST" {
// 			arrCopy = append(arrCopy[:i-1], arrCopy[i+1:]...)
// 		}
// 		if arrCopy[i-1] == "WEST" && arrCopy[i] == "EAST" {
// 			arrCopy = append(arrCopy[:i-1], arrCopy[i+1:]...)
// 		}
// 	}
// 	return arrCopy
// }

func main() {
	// var a = []string{"NORTH", "SOUTH", "EAST", "WEST"}
	var b = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
	// fmt.Println(DirReduc(a))
	fmt.Println(DirReduc(b))
}
