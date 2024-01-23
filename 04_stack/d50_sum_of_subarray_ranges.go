package main

import "fmt"

func subArrayRanges(nums []int) int64 {
	var stack []int
	var prevMinIndex int
	var res int
	for right := 0; right <= len(nums); right++ {
		for len(stack) > 0 && (right == len(nums) || nums[stack[len(stack)-1]] >= nums[right]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				prevMinIndex = stack[len(stack)-1]
			} else {
				prevMinIndex = -1
			}
			res -= nums[top] * (top - prevMinIndex) * (right - top)
		}
		stack = append(stack, right)
	}
	stack = []int{}
	for right := 0; right <= len(nums); right++ {
		for len(stack) > 0 && (right == len(nums) || nums[stack[len(stack)-1]] <= nums[right]) {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				prevMinIndex = stack[len(stack)-1]
			} else {
				prevMinIndex = -1
			}
			res += nums[top] * (top - prevMinIndex) * (right - top)
		}
		stack = append(stack, right)
	}
	ans := int64(res)
	return ans
}

func main() {
	//testArray := []int{1, 3, 3} // -> 4
	testArray := []int{4, -2, -3, 4, 1} // -> 59
	fmt.Println(subArrayRanges(testArray))
}
