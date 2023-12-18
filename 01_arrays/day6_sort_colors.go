// https://leetcode.com/problems/sort-colors/

package main

import "fmt"

func sortColors(nums []int) []int {
	left, right := 0, len(nums)-1
	i := 0
	for i <= right {
		if nums[i] == 2 {
			nums[right], nums[i] = nums[i], nums[right]
			right--
			if nums[i] == 1 {
				i++
			}
		} else if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			i++
			left++
		} else {
			i++
		}
	}
	return nums
}

func main() {
	testNums := []int{2, 0, 2, 1, 1, 0}
	fmt.Println(sortColors(testNums))
}
