package main

import (
	"fmt"
)

func removeDuplicates(s string) string {
	var stack []rune
	for _, char := range s {
		if len(stack) == 0 || char != stack[len(stack)-1] {
			stack = append(stack, char)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	res := string(stack)
	return res
}

func main() {
	testString := "abbaca"
	fmt.Println(removeDuplicates(testString))
}
