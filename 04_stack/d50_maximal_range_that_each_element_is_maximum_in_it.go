// https://leetcode.com/problems/maximal-range-that-each-element-is-maximum-in-it/description/

package main

import "fmt"

func maximumLengthOfRanges(nums []int) []int {
	var stack []int
	res := make([]int, len(nums))
	var rightBiggest int
	for idx, num := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[idx] = stack[len(stack)-1] + 1
		} else {
			res[idx] = 0
		}
		stack = append(stack, idx)
	}
	stack = []int{}
	for right := len(nums) - 1; right >= 0; right-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[right] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			rightBiggest = stack[len(stack)-1] - 1
		} else {
			rightBiggest = len(nums) - 1
		}
		res[right] = rightBiggest - res[right] + 1
		stack = append(stack, right)
	}
	return res
}

func main() {
	testNums := []int{1, 5, 4, 3, 6}
	fmt.Println(maximumLengthOfRanges(testNums))
}
