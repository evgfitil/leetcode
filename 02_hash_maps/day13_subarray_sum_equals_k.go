// https://leetcode.com/problems/subarray-sum-equals-k/description/

package main

import "fmt"

func subarraySum(nums []int, k int) int {
	res, currentSum := 0, 0
	prefixSums := map[int]int{0: 1}
	for _, num := range nums {
		currentSum += num
		prefix := currentSum - k
		res += prefixSums[prefix]
		prefixSums[currentSum] += 1
	}
	return res
}

func main() {
	testNumsOne := []int{1, 1, 1}
	kOne := 2
	testNumsTwo := []int{1, 2, 3}
	kTwo := 3
	fmt.Println(subarraySum(testNumsOne, kOne))
	fmt.Println(subarraySum(testNumsTwo, kTwo))
}
