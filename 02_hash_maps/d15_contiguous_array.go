// https://leetcode.com/problems/contiguous-array/description/

package main

import "fmt"

func findMaxLength(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	indexes := map[int]int{0: -1}
	balance := 0
	for idx, num := range nums {
		balance += num
		if _, ok := indexes[balance]; ok {
			res = max(res, idx-indexes[balance])
		} else {
			indexes[balance] = idx
		}
	}
	return res
}

func main() {
	testArray := []int{0, 1, 1, 0, 0, 1, 1, 0}
	fmt.Println(findMaxLength(testArray))
}
