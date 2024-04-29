package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Brightest(colors []string) string {
	index := 0
	for i, color := range colors {
		var brightest int64
		decimalNum, err := strconv.ParseInt(strings.TrimPrefix(color, "#"), 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		if decimalNum < brightest {
			brightest = decimalNum
			index = i
		}
		fmt.Println(i, decimalNum)
	}
	return colors[index]
}
func main() {
	// colors := []string{"#A66B87", "#BA482F", "#72B64E", "#E97FFE", "#E3D804", "#04C46D", "#3A8421", "#3ED2A1", "#9C9427", "#7EE19C", "#AE2546", "#F4BD1B", "#1790F4", "#F16C8B", "#74E14C"}
	colors := []string{"#1F", "#6"}
	fmt.Println(Brightest(colors))
}
