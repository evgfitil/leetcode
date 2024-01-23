// https://leetcode.com/problems/find-the-most-competitive-subsequence/description/

package main

import "fmt"

func mostCompetitive(nums []int, k int) []int {
	var stack []int

	for i := 0; i < len(nums); i++ {
		curNum := nums[i]
		for len(stack) > 0 && stack[len(stack)-1] > curNum && len(stack)-1+len(nums)-i >= k {
			stack = stack[:len(stack)-1]
		}
		if len(stack) < k {
			stack = append(stack, nums[i])
		}
	}
	return stack
}

func main() {
	//testNums := []int{2, 4, 3, 3, 5, 4, 9, 6}
	//k := 4
	testNums := []int{71, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2} // -> [8,80,2]
	k := 3
	fmt.Println(mostCompetitive(testNums, k))
}
