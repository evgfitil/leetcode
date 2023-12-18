package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if abs(nums[left]) < abs(nums[right]) {
			res[i] = nums[right] * nums[right]
			right--
		} else {
			res[i] = nums[left] * nums[left]
			left++
		}
	}

	return res
}

func main() {
	testNums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(testNums))
}
