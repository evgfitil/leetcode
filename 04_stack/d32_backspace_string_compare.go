package main

import "fmt"

func cutTheString(s string) []rune {
	var stack []rune
	for _, char := range s {
		switch {
		case len(stack) > 0 && char == rune('#'):
			stack = stack[:len(stack)-1]
		case char != rune('#'):
			stack = append(stack, char)
		}
	}
	return stack
}
func backspaceCompare(s string, t string) bool {
	first, second := cutTheString(s), cutTheString(t)
	if len(first) != len(second) {
		return false
	}
	for i := len(first) - 1; i >= 0; i-- {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func main() {
	//testS, testT := "ab#c", "ad#c"
	testS, testT := "y#fo##f", "y#f#o##f"
	fmt.Println(backspaceCompare(testS, testT))
}
