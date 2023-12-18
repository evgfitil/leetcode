// https://leetcode.com/problems/find-lucky-integer-in-an-array/description/

package main

import "fmt"

func findLucky(arr []int) int {
	count := make(map[int]int)
	for _, v := range arr {
		count[v] = count[v] + 1
	}
	max := -1
	for k, v := range count {
		if k == v {
			if max < k {
				max = k
			}
		}
	}
	return max
}

func main() {
	testNums := []int{1, 2, 2, 3, 3, 3}
	fmt.Println(findLucky(testNums))
}
