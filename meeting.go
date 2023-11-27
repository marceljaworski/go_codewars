package main

import (
	"fmt"
	"sort"
	"strings"
)

type Friend struct {
	lastName  string
	firstName string
}

func Meeting(s string) string {
	s = strings.ToUpper(s)
	friendsList := strings.Split(s, ";")
	sort.Strings(friendsList)
	sortedList := []Friend{}
	for _, name := range friendsList {
		names := strings.Split(name, ":")
		v := Friend{names[1], names[0]}
		sortedList = append(sortedList, v)
	}
	sort.SliceStable(sortedList, func(i, j int) bool {
		return sortedList[i].lastName < sortedList[j].lastName
	})
	solution := ""
	for _, guest := range sortedList {
		solution += fmt.Sprintf("(%s, %s)", guest.lastName, guest.firstName)
	}
	return solution
}
func main() {
	fmt.Println(Meeting("Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn"))
}
