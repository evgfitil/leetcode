// https://leetcode.com/problems/reverse-string/

package main

import "fmt"

func reverseString(s []byte) {
	fmt.Println(recursiveSwap(0, len(s)-1, s))
}

func recursiveSwap(left, right int, s []byte) []byte {
	if left >= right {
		return s
	}
	s[left], s[right] = s[right], s[left]
	return recursiveSwap(left+1, right-1, s)
}

func main() {
	testInput := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(testInput)
}
