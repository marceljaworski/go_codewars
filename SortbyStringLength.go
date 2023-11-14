package main

import (
	"fmt"
	"sort"
)

//	func SortByLength(arr []string) []string {
//		for i := 0; i < len(arr); i++ {
//			for j := i; j < len(arr); j++ {
//				fmt.Println(i, j)
//				if len(arr[j]) < len(arr[i]) {
//					arr[j], arr[i] = arr[i], arr[j]
//				}
//			}
//		}
//		return arr
//	}
func SortByLength(a []string) []string {
	sort.Slice(a, func(i, j int) bool { return len(a[i]) < len(a[j]) })
	return a
}
func main() {
	fmt.Println(SortByLength([]string{"Telescopes", "Glasses", "Eyes", "Monocles"}))
}
