// https://leetcode.com/problems/using-a-robot-to-print-the-lexicographically-smallest-string/description/

package main

import "fmt"

func robotWithString(s string) string {
	var p, t []rune
	minCharArr := make([]rune, len(s))
	minChar := rune(137)
	for i := len(s) - 1; i >= 0; i-- {
		if rune(s[i]) < minChar {
			minChar = rune(s[i])
		}
		minCharArr[i] = minChar
	}

	for i, char := range s {
		for len(t) > 0 && t[len(t)-1] <= minCharArr[i] {
			top := t[len(t)-1]
			p = append(p, top)
			t = t[:len(t)-1]
		}
		t = append(t, char)
	}
	for len(t) > 0 {
		top := t[len(t)-1]
		t = t[:len(t)-1]
		p = append(p, top)
	}
	return string(p)
}

func main() {
	//testS := "zza"
	//testS := "bac"
	//testS := "bdda"
	testS := "abcabc"
	fmt.Println(robotWithString(testS))
}
