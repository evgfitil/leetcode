// https://leetcode.com/problems/next-greater-element-i/description/

package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	nextGreater := make(map[int]int)
	for _, num := range nums2 {
		nextGreater[num] = -1
	}
	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGreater[top] = num
		}
		stack = append(stack, num)
	}
	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = nextGreater[num]
	}
	return res
}

func main() {
	testNums1 := []int{4, 1, 2}
	testNums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(testNums1, testNums2))
}
