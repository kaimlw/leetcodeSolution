package main

import "fmt"

func mergeAlternately(word1, word2 string) string {
	word_length := 0
	result := ""

	if len(word1) >= len(word2) {
		word_length = len(word1)
	} else {
		word_length = len(word2)
	}

	for i := 0; i < word_length; i++ {
		if i < len(word1) {
			result += string(word1[i])
		}
		if i < len(word2) {
			result += string(word2[i])
		}
	}

	return result
}

func main() {
	word1 := "abcd"
	word2 := "pq"

	fmt.Println(mergeAlternately(word1,word2))
}