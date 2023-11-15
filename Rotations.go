package main

import "fmt"

func ContainAllRots(strng string, arr []string) bool {
	counter := 0
	master := []string{}
	//rotate string and append to master
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
