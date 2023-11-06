package main

import (
	"fmt"
	"strings"
)

//	func VertMirror(s string) string {
//		slice := strings.Split(s, "\n")
//		sol := make([]string, 0, len(slice))
//		for _, word := range slice {
//			wordSlice := make([]string, 0, len(word))
//			for _, char := range word {
//				letter := string(char)
//				wordSlice = append([]string{letter}, wordSlice...)
//			}
//			reversedWord := "..."
//			reversedWord += strings.Join(wordSlice, "")
//			// sol = append(sol, reversedWord)
//			sol = append([]string{reversedWord}, sol...)
//		}
//		return strings.Join(sol, "\n")
//	}
//
//	func HorMirror(s string) string {
//		slice := strings.Split(s, "\n")
//		sol := make([]string, 0, len(slice))
//		for _, word := range slice {
//			// sol = append([]string{word}, sol...)
//			word += "..."
//			sol = append(sol, word)
//		}
//		return strings.Join(sol, "\n")
//	}
func Rot(s string) string {
	slice := strings.Split(s, "\n")
	sol := make([]string, 0, len(slice))
	for _, word := range slice {
		wordSlice := make([]string, 0, len(word))
		for _, char := range word {
			letter := string(char)
			wordSlice = append([]string{letter}, wordSlice...)
		}
		reversedWord := strings.Join(wordSlice, "")
		sol = append([]string{reversedWord}, sol...)
	}
	return strings.Join(sol, "\n")
}

func SelfieAndRot(s string) string {
	slice := strings.Split(s, "\n")
	rot := make([]string, 0, len(slice))
	sol := make([]string, 0, len(slice))
	//selfie
	for _, word := range slice {
		sol = append(sol, word)
	}
	//rot
	for _, word := range slice {
		wordSlice := make([]string, 0, len(word))
		for _, char := range word {
			letter := string(char)
			wordSlice = append([]string{letter}, wordSlice...)
		}
		reversedWord := strings.Join(wordSlice, "")
		rot = append([]string{reversedWord}, rot...)
	}
	dots := ""
	for len(dots) < len(slice) {
		dots += "."
	}
	dots += "\n"
	sol = append(sol, rot...)
	return strings.Join(sol, dots)
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

//	func main() {
//		fmt.Println("___Vertical Mirror___\n", Oper(VertMirror, "abcd\nefgh\nijkl\nmnop"))
//		fmt.Println("___Horizpntal Mirror___\n", Oper(HorMirror, "abcd\nefgh\nijkl\nmnop"))
//	}
func main() {
	fmt.Println("___Selfie and Rot___\n", Oper(SelfieAndRot, "xZBVf\njsbSf\nJcpNf\nfVnPf\nabcde"))
	fmt.Println("___Rot___\n", Oper(Rot, "abcd\nefgh\nijkl\nmnop"))
}
