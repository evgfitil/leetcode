package main

import "fmt"

func validSubarrays(nums []int) int {
	var stack []int
	res := 0
	for idx, value := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > value {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res += idx - top
		}
		stack = append(stack, idx)
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res += len(nums) - top
	}
	return res + len(stack)
}

func main() {
	testArray := []int{1, 4, 2, 5, 3}
	fmt.Println(validSubarrays(testArray))
}
