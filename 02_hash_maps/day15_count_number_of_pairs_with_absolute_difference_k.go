// https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/description/

package main

import "fmt"

func countKDifference(nums []int, k int) int {
	res := 0
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num] += 1
	}
	for num, count := range counts {
		if counts[num+k] > 0 {
			res += count * counts[num+k]
		}
	}
	return res
}

func main() {
	testNums := []int{1, 2, 2, 1}
	k := 1
	fmt.Println(countKDifference(testNums, k))
}
