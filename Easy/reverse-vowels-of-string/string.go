package main

import (
	"fmt"
	"regexp"
	"strings"
)

func reverseVowels(s string) string {
	regex := regexp.MustCompile("[aiueoAIUEO]")
	splittedString := strings.Split(s, "")
	vowels := []string{}

	for _,item := range splittedString {
		if regex.MatchString(string(item)) {
			vowels = append(vowels, item)
		}
	}

	idx := len(vowels)-1
	for i, item := range splittedString{
		if regex.MatchString(string(item)) {
			splittedString[i] = vowels[idx]
			idx--
		}
	}

	return strings.Join(splittedString,"")
}

func reverseVowels2(s string) string {
	found := []byte(s)
	i := 0
	j := len(s) - 1

	for i < j {
			if !isVowel(s[i]) {
					i++
					continue
			}
			if !isVowel(s[j]) {
					j--
					continue
			}
			found[i], found[j] = found[j], found[i]
			i++
			j--
	}
	return string(found)
}

func isVowel(c byte) bool {
	vowels := []byte{'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u'}
	for _, v := range vowels {
			if c == v {
					return true
			}
	}
	return false
}

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
	fmt.Println(reverseVowels2("IceCreAm"))
	
}