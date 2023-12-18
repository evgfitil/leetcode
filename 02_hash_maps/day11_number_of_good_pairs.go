// https://leetcode.com/problems/number-of-good-pairs/description/

package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	pairs := make(map[int]int)
	pairs_increment := make(map[int]int)
	for _, v := range nums {
		if pairs_increment[v] > 0 {
			pairs[v] = pairs[v] + pairs_increment[v]
		}
		pairs_increment[v] = pairs_increment[v] + 1
	}
	sum := 0
	for _, k := range pairs {
		sum = sum + k
	}
	return sum
}

// func numIdenticalPairs(nums []int) int {
// 	pairs := make(map[int]int)
// 	res := 0
// 	for _, v := range nums {
// 		res = res + pairs[v]
// 		pairs[v] = pairs[v] + 1
// 	}
// 	return res
// }

func main() {
	testNums := []int{1, 2, 3, 1, 1, 3}
	// testNums := []int{1, 1, 1, 1}
	fmt.Println(numIdenticalPairs(testNums))
}
