package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	res := getRow(rowIndex - 1)
	lastRow := res
	temp := make([]int, len(lastRow)+2)
	copy(temp[1:], lastRow)
	tempRes := make([]int, rowIndex+1)
	for j := 0; j <= len(tempRes)-1; j++ {
		tempRes[j] = temp[j] + temp[j+1]
	}
	res = tempRes
	return res
}

func main() {
	testN := 3
	fmt.Println(getRow(testN))
}
