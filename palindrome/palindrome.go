package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {

		for i < j && !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i++
		}

		for i < j && !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
			j--
		}

		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("radar")) // true
	fmt.Println(isPalindrome("hello")) // false
}
