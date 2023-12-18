// https://leetcode.com/problems/word-pattern/description/

package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	charsToWords := make(map[rune]string)
	wordsToChars := make(map[string]rune)

	for i, char := range pattern {
		word := words[i]

		mappedWord, ok := charsToWords[char]
		if ok {
			if mappedWord != word {
				return false
			}
		} else {
			charsToWords[char] = word
		}

		mappedChar, ok := wordsToChars[word]
		if ok {
			if mappedChar != char {
				return false
			}
		} else {
			wordsToChars[word] = char
		}
	}
	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat fish"
	fmt.Println(wordPattern(pattern, s))
}
