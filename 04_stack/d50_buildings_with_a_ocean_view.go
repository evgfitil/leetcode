package main

import "fmt"

func findBuildings(heights []int) []int {
	var stack []int
	for right := 0; right < len(heights); right++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] <= heights[right] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, right)
	}
	return stack
}

func main() {
	testBuildings := []int{4, 2, 3, 1} // -> 0 2 3
	//testBuildings := []int{4, 3, 2, 1} // -> 0 1 2 3
	//testBuildings := []int{1, 3, 2, 4} // -> 3
	fmt.Println(findBuildings(testBuildings))
}
