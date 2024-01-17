package main

import "fmt"

func sumSubarrayMins(arr []int) int {
	const mod = 1000000007
	var stack []int
	var prevMinIndex int
	res := 0
	for idx, value := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= value {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				prevMinIndex = stack[len(stack)-1]
			} else {
				prevMinIndex = -1
			}
			res += arr[top] * (top - prevMinIndex) * (idx - top)
			res %= mod
		}
		stack = append(stack, idx)
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			prevMinIndex = stack[len(stack)-1]
		} else {
			prevMinIndex = -1
		}
		res += arr[top] * (top - prevMinIndex) * (len(arr) - top)
		res %= mod
	}
	return res
}

func main() {
	testArray := []int{3, 1, 2, 4}
	fmt.Println(sumSubarrayMins(testArray))
}
