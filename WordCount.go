package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	// phrase = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(phrase, ".", ""), ",", ""), `"`, ""), "?", ""), ":", ""), "!", "")
	// slice := strings.Fields(strings.ToLower(phrase))
	re := regexp.MustCompile(`\w+('\w+)?`)
	words := make(Frequency)
	for _, w := range re.FindAllString(strings.ToLower(phrase), -1) {
		words[w]++
	}
	// for _, word := range slice {
	// 	if after, found := strings.CutPrefix(word, "'"); found {
	// 		word = after
	// 	}
	// 	if after, found := strings.CutSuffix(word, "'"); found {
	// 		word = after
	// 	}
	// 	words[word]++
	// }

	return words
}
func main() {
	fmt.Println(WordCount(`"That's the password: 'PASSWORD 123'!", cried the Special Agent. So I fled.`))
}
