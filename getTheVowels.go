package main

func GetTheVowels(word string) int {
	result := 0
	index := 0
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, char := range word {
		if string(char) == vowels[index] {
			result++
			if index == 4 {
				index = 0
			} else {
				index++
			}
		}
	}
	return result
}
