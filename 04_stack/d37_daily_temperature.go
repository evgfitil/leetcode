package main

import "fmt"

//// time limit bruteforce
//func dailyTemperatures(temperatures []int) []int {
//	res := make([]int, len(temperatures))
//	right := 0
//	for left := 0; left < len(temperatures)-1; left++ {
//		count := 0
//		for right < len(temperatures)-1 && temperatures[right] <= temperatures[left] {
//			count++
//			right++
//		}
//		if temperatures[left] < temperatures[right] {
//			res[left] = count
//		}
//		right = left + 1
//	}
//	return res
//}

func dailyTemperatures(temperatures []int) []int {
	var stack []int
	res := make([]int, len(temperatures))
	for idx, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] = idx - top
		}
		stack = append(stack, idx)
	}
	return res
}

func main() {
	testNums := []int{73, 74, 75, 71, 69, 72, 76, 73}
	//testNums := []int{30, 40, 50, 60}
	fmt.Println(dailyTemperatures(testNums))
}
