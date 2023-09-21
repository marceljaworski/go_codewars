package main

import (
	"fmt"
	"time"
)

func UnluckyDays(year int) int {
	tot := 0
	for m := 1; m <= 12; m++ {
		if time.Date(year, time.Month(m), 13, 0, 0, 0, 0, time.UTC).Weekday() == 5 {
			tot++
		}
	}

	return tot
}

func main() {

	// start := time.Now()
	// t := time.Now()
	// elapsed := t.Sub(start)
	// fmt.Println(start, t, elapsed)
	// loc, _ := time.LoadLocation("Europe/Berlin")

	// const shortForm = "2006-Jan-02"
	// t :=
	// t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	// fmt.Println(t)

	fmt.Println(UnluckyDays(2024))
}
