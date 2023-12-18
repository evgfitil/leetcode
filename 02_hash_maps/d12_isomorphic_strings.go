// https://leetcode.com/problems/isomorphic-strings/description/

package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	mapS := make(map[rune]rune)
	mapT := make(map[rune]rune)
	for i := 0; i < len(s); i++ {
		if mapS[rune(s[i])] != 0 && mapS[rune(s[i])] != rune(t[i]) ||
			mapT[rune(t[i])] != 0 && mapT[rune(t[i])] != rune(s[i]) {
			return false
		}
		mapS[rune(s[i])] = rune(t[i])
		mapT[rune(t[i])] = rune(s[i])
	}
	return true
}

func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))
}
