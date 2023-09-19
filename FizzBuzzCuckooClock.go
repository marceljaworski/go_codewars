package main

import (
	"strconv"
	"strings"
)

func FizzBuzzCuckooClock(time string) string {

	timeNow := strings.Split(time, ":")
	hour, err := strconv.Atoi(timeNow[0])
	if err != nil {
		panic(err)
	}
	minute, err := strconv.Atoi(timeNow[1])
	if err != nil {
		panic(err)
	}
	if hour > 12 {
		hour = hour - 12
	}

	var results []string

	if minute%3 == 0 && minute > 0 && minute != 30 {
		results = append(results, "Fizz")
	}
	if minute%3 != 0 && minute%5 != 0 {
		results = append(results, "tick")
	}
	if minute%5 == 0 && minute > 0 && minute != 30 {
		results = append(results, "Buzz")
	} else if minute == 30 {
		results = append(results, "Cuckoo")
	} else if minute == 0 && hour > 0 {

		for i := 0; i < hour; i++ {
			results = append(results, "Cuckoo")
		}
	} else if minute == 0 && hour == 0 {

		for i := 0; i < 12; i++ {
			results = append(results, "Cuckoo")
		}
	}
	result := strings.Join(results, " ")
	return result
}
