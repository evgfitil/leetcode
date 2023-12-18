// https://leetcode.com/problems/decode-the-message/description/

package main

import (
	"fmt"
	"strings"
)

func decodeMessage(key string, message string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetIdx := 0
	alphabetMap := make(map[rune]rune)
	keySeen := make(map[rune]bool)
	for _, char := range key {
		if !keySeen[char] && char >= 'a' && char <= 'z' {
			keySeen[char] = true
			alphabetMap[char] = rune(alphabet[alphabetIdx])
			alphabetIdx++
		}
	}
	var decodedMessage strings.Builder
	for _, char := range message {
		if decodedChar, ok := alphabetMap[char]; ok {
			decodedMessage.WriteRune(decodedChar)
		} else {
			decodedMessage.WriteRune(char)
		}
	}
	return decodedMessage.String()
}

func main() {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	fmt.Println(decodeMessage(key, message))
}
