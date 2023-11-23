package main

import "fmt"

func PrinterError(s string) string {
	errorCount := 0
	lenght := len(s)
	for _, char := range s {
		if char >= 110 && char < 123 {
			errorCount++
		}
	}
	return fmt.Sprintf("%d/%d", errorCount, lenght)
}
func main() {
	fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
}
