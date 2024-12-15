package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	splittedString := strings.Split(strings.TrimSpace(s)," ")
	cleanString := []string{}
	for _,item := range splittedString {
		if len(item) == 0 {
			continue
		}
		cleanString = append(cleanString, item)
	}

	first := 0
	last := len(cleanString)-1
	for first < last {
		cleanString[first], cleanString[last] = cleanString[last], cleanString[first]
		first++
		last--
	}

	return strings.Join(cleanString, " ")
}

func reverseWords2(s string) string {
	words := strings.Fields(s) 

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
			words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	// fmt.Println(reverseWords("the sky is blue"))
	// fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWords("a good   example"))
	fmt.Println(reverseWords2("a good   example"))
}