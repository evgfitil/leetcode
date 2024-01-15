// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/description/

package main

import "fmt"

func finalPrices(prices []int) []int {
	var stack []int
	res := make([]int, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {
		discount := 0
		for len(stack) > 0 && stack[len(stack)-1] > prices[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			discount = stack[len(stack)-1]
		}
		res[i] = prices[i] - discount
		stack = append(stack, prices[i])
	}
	return res
}

func main() {
	//testPrices := []int{8, 4, 6, 2, 3} // -> 4 2 4 2 3
	testPrices := []int{10, 1, 1, 6} // -> 9 0 1 6
	fmt.Println(finalPrices(testPrices))
}
