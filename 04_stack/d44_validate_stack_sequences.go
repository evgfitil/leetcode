// https://leetcode.com/problems/validate-stack-sequences/description/

package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	left, right := 0, 0
	for right < len(popped) {
		for left < len(pushed) && (len(stack) == 0 || stack[len(stack)-1] != popped[right]) {
			stack = append(stack, pushed[left])
			left++
		}
		if len(stack) > 0 && stack[len(stack)-1] != popped[right] {
			return false
		}
		stack = stack[:len(stack)-1]
		right++
	}
	return true
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	//pushed := []int{1, 0}
	//popped := []int{1, 0}
	//pushed := []int{1, 2, 3, 4, 5}
	//popped := []int{4, 3, 5, 1, 2}
	fmt.Println(validateStackSequences(pushed, popped))
}
