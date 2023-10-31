package main

import "fmt"

func ContainAllRots(strng string, arr []string) bool {
	// solution := 0
	// if solution == len(strng) {
	// 	return true
	// }
	// for _, char := range strng {
	// 	count := len(strng)*2 - 2
	// 	if count == 0 {
	// 		solution++
	// 	}
	// 	for _, word := range arr {
	// 		for i, letter := range word {
	// 			if letter == char {
	// 				count = count - i
	// 			}
	// 		}
	// 	}
	// }
	// return false

	counter := 0
	master := []string{}
	for i, _ := range strng {
		master = append(master, strng[i:]+strng[:i])
	}
	for _, j := range master {
		for _, k := range arr {
			if j == k {
				counter++
				break
			}
		}
	}
	return len(strng) == counter

}
func main() {
	fmt.Println(ContainAllRots("bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}))
}
