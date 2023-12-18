package leetcoders

import "fmt"

func findNonMinOrMax(nums []int) int {
	res := -1
	if len(nums) <= 2 {
		return res
	}
	min, max := nums[0], nums[0]
	for _, n := range nums {
		switch {
		case n < max && n > min:
			return n
		case n > max:
			res = max
			max = n
		case n < min:
			res = min
			min = n
		}
	}
	return res
}

func main() {
	testNums := []int{2, 1, 3}
	fmt.Println(findNonMinOrMax(testNums))
}
