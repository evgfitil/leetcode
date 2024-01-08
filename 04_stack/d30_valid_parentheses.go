// https://leetcode.com/problems/valid-parentheses/description/

package main

import "fmt"

func isValid(s string) bool {
	var stack []rune
	braces := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, brace := range s {
		if closingBrace, exists := braces[brace]; exists {
			stack = append(stack, closingBrace)
		} else {
			if len(stack) == 0 {
				return false
			}
			if brace != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	testString := "()[]{}"
	fmt.Println(isValid(testString))
}
