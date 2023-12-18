// https://leetcode.com/problems/contains-duplicate/

package main

import "fmt"

func containsDuplicate(nums []int) bool {
	temp := make(map[int]int)
	for _, v := range nums {
		if temp[v] >= 1 {
			return true
		}
		temp[v] = temp[v] + 1
	}
	return false
}

func main() {
	testNums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(testNums))
}
