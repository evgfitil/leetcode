// https://leetcode.com/problems/binary-subarrays-with-sum/description/

package main

import "fmt"

func numSubarraysWithSum(nums []int, goal int) int {
	count, currentSum := 0, 0
	sums := make(map[int]int)
	sums[0] = 1
	for _, num := range nums {
		currentSum = currentSum + num
		count = count + sums[currentSum-goal]
		sums[currentSum] = sums[currentSum] + 1
	}
	return count
}

func main() {
	testNumsOne := []int{1, 0, 1, 0, 1}
	goalOne := 2
	testNumsTwo := []int{0, 0, 0, 0, 0}
	goalTwo := 0
	fmt.Println(numSubarraysWithSum(testNumsOne, goalOne))
	fmt.Println(numSubarraysWithSum(testNumsTwo, goalTwo))
}
