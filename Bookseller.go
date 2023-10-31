package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
	myMap := map[string]int{}
	for _, category := range listCat {
		myMap[category] = 0
		for _, word := range listArt {
			code := strings.Fields(word)
			if code[0][0:1] == string(category) {
				i, _ := strconv.Atoi(code[1])
				myMap[category] += i
			}
		}
	}
	slice := make([]string, 0, len(myMap))
	for _, v := range listCat {
		slice = append(slice, fmt.Sprintf("(%v : %v)", v, myMap[v]))
	}

	return strings.Join(slice, " - ")
}
func main() {
	a := []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}
	b := []string{"A", "B"}
	var c = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var d = []string{"A", "B", "C", "D"}
	fmt.Println(StockList(a, b))
	fmt.Println(StockList(c, d))
}
