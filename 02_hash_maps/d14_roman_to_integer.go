// https://leetcode.com/problems/roman-to-integer/description/

package main

import "fmt"

func romanToInt(s string) int {
	roman := map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}
	res, previous := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		current := roman[rune(s[i])]
		if current < previous {
			res -= current
		} else {
			res += current
		}
		previous = current
	}
	return res
}

func main() {
	testString := "III"
	fmt.Println(romanToInt(testString))
}
