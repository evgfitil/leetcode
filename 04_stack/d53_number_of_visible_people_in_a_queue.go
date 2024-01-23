// https://leetcode.com/problems/number-of-visible-people-in-a-queue/description/

package main

import "fmt"

func canSeePersonsCount(heights []int) []int {
	var st []int
	res := make([]int, len(heights))
	for right := 0; right <= len(heights)-1; right++ {
		for len(st) > 0 && heights[st[len(st)-1]] < heights[right] {
			top := st[len(st)-1]
			st = st[:len(st)-1]
			res[top]++
		}
		if len(st) > 0 {
			res[st[len(st)-1]]++
		}
		st = append(st, right)
	}
	return res
}

func main() {
	testHeights := []int{10, 6, 8, 5, 11, 9}
	fmt.Println(canSeePersonsCount(testHeights))
}
