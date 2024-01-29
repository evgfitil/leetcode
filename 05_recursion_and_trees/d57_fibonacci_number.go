package main

import "fmt"

func fib(n int, memo map[int]int) int {
	//// recursive
	//if n <= 1 {
	//	return n
	//}
	//return fib(n-1) + fib(n-2)

	// memo
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fib(n-1, memo) + fib(n-2, memo)

	return memo[n]

	//// iterative
	//a, b := 0, 1
	//for i := 0; i < n; i++ {
	//	a, b = b, b+a
	//}
	//return a

}

func main() {
	memo := make(map[int]int)
	testN := 4
	fmt.Println(fib(testN, memo))
}
