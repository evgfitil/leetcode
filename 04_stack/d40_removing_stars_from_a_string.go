// https://leetcode.com/problems/removing-stars-from-a-string/

package main

import (
	"fmt"
	"strings"
)

type stack struct {
	elements []string
}

func (s *stack) Push(val string) {
	s.elements = append(s.elements, val)
}

func (s *stack) Pop() {
	s.elements = s.elements[:len(s.elements)-1]
}

func removeStars(s string) string {
	var st stack
	for _, value := range strings.Split(s, "") {
		if value == "*" && len(st.elements) > 0 {
			st.Pop()
		} else {
			st.Push(value)
		}
	}
	return strings.Join(st.elements, "")
}

func main() {
	testString := "leet**cod*e"
	fmt.Println(removeStars(testString))
}
