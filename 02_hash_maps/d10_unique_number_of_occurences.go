// https://leetcode.com/problems/unique-number-of-occurrences/description/

package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)
	for _, num := range arr {
		count[num] = count[num] + 1
	}
	set := make(map[int]bool)
	for _, v := range count {
		if _, ok := set[v]; !ok {
			set[v] = true
		} else {
			return false
		}
	}
	return true
}

func main() {
	// testNums := []int{1, 2, 2, 1, 1, 3}
	testNums := []int{1, 2}
	fmt.Println(uniqueOccurrences(testNums))
}
