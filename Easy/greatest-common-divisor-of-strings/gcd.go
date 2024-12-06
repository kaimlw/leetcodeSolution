package main

import (
	"fmt"
	"strings"
)

func gcd(num1, num2 int) int {
	for num1 > 0 && num2 > 0 {
		if num1 > num2 {
			num1 = num1 % num2
		} else {
			num2 = num2 % num1
		}
	}

	if num1 > 0 {
		return num1
	} else {
		return num2
	}
}

func gcdOfStrings(str1 string, str2 string) string {
	gcdInt := gcd(len(str1), len(str2))

	if str1[:gcdInt] == str2[:gcdInt] {
		gcdStr := str1[:gcdInt]
		if len(strings.ReplaceAll(str1, gcdStr, "")) != 0 {
			return ""
		}
		if len(strings.ReplaceAll(str2, gcdStr, "")) != 0 {
			return ""
		}

		return str1[:gcdInt]
	}

	return ""
}

// RUNTIME 0ms Solution
func gcdOfStrings2(str1 string, str2 string) string {
    if str1 == "" {
        return str2
    }
    if str2 == "" {
        return str1
    }
    if len(str1) > len(str2) {
        trimmedStr1 := strings.TrimSuffix(str1, str2)
        if len(trimmedStr1) == len(str1) {
            return ""
        }
        return gcdOfStrings2(trimmedStr1, str2)
    }
    trimmedStr2 := strings.TrimSuffix(str2, str1)
    if len(trimmedStr2) == len(str2) {
        return ""
    }
    return gcdOfStrings2(str1, trimmedStr2)
}
// -------------------

// LEAST MEMORY (4.26 MB) SOLUTION
func gcdOfStrings3(str1 string, str2 string) string {
	least := str2
	if len(str1) < len(str2) {
		least = str1
	}

	for i := len(least); i > 0; i-- {
		str := least[0:i]
		if len(str1)%i != 0 || len(str2)%i != 0 {
			continue
		}
		if dividable(str1, str) && dividable(str2, str) {
			return str
		}
	}
	return ""
}

func dividable(s, t string) bool {
	// assume divisible
	i := 0
	j := len(t)
	for j <= len(s) {
		if s[i:j] != t {
			return false
		}
		i = j
		j = j + len(t)
	}
	return true
}
// ---------------------------------

func main() {
	fmt.Println(gcdOfStrings("ABABCCABAB", "ABAB"))
	fmt.Println(gcdOfStrings("ABCABD", "ABC"))
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("TAUXXTAUXXTAUXXTAUXXTAUXX", "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"))
	fmt.Println(gcdOfStrings("ABCDEF", "ABC"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))
	fmt.Println(gcdOfStrings("AAAAAAAAA", "AAACCC"))
	// fmt.Println(gcd(48,18))
}
