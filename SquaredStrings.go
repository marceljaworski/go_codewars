package main

import (
	"fmt"
	"strings"
)

func VertMirror(s string) string {
	slice := strings.Split(s, "\n")
	sol := make([]string, 0, len(slice))
	for _, word := range slice {
		wordSlice := make([]string, 0, len(word))
		for _, char := range word {
			letter := string(char)
			wordSlice = append([]string{letter}, wordSlice...)
		}
		reversedWord := strings.Join(wordSlice, "")
		sol = append(sol, reversedWord)
	}
	return strings.Join(sol, "\n")
}
func HorMirror(s string) string {
	slice := strings.Split(s, "\n")
	sol := make([]string, 0, len(slice))
	for _, word := range slice {
		sol = append([]string{word}, sol...)
	}
	return strings.Join(sol, "\n")
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}
func main() {
	fmt.Println("___Vertical Mirror___\n", Oper(VertMirror, "abcd\nefgh\nijkl\nmnop"))
	fmt.Println("___Horizpntal Mirror___\n", Oper(HorMirror, "abcd\nefgh\nijkl\nmnop"))
}
