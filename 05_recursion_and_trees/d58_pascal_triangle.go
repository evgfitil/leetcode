package main

import "fmt"

//// iterative
//
//func generate(numRows int) [][]int {
//	res := [][]int{{1}}
//	for i := 1; i <= numRows-1; i++ {
//		lastRow := res[len(res)-1]
//		temp := make([]int, len(lastRow)+2)
//		copy(temp[1:], lastRow)
//		tempRes := make([]int, i+1)
//		for j := 0; j <= len(tempRes)-1; j++ {
//			tempRes[j] = temp[j] + temp[j+1]
//		}
//		res = append(res, tempRes)
//	}
//	return res
//}

// recursive
func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	res := generate(numRows - 1)
	lastRow := res[len(res)-1]
	temp := make([]int, len(lastRow)+2)
	copy(temp[1:], lastRow)
	tempRes := make([]int, numRows)
	for j := 0; j <= len(tempRes)-1; j++ {
		tempRes[j] = temp[j] + temp[j+1]
	}
	res = append(res, tempRes)

	return res
}

func main() {
	testN := 5
	fmt.Println(generate(testN))
}
