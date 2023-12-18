// https://leetcode.com/problems/calculate-money-in-leetcode-bank/description/

package main

import "fmt"

func totalMoney(n int) int {
	weekNumber := 0
	sum := 0
	dailyIncrement := 0
	for i := 1; i <= n; i++ {
		if i%7 == 1 {
			weekNumber++
		}
		sum = sum + weekNumber + dailyIncrement
		dailyIncrement++
	}
	return sum
}

func main() {
	n := 4
	fmt.Println(totalMoney(n))
}
