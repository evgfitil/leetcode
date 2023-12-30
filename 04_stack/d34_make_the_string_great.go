package main

import (
	"errors"
	"fmt"
	"unicode"
)

type Stack struct {
	values []rune
}

func (s *Stack) Push(char rune) {
	s.values = append(s.values, char)
}

func (s *Stack) Length() int {
	return len(s.values)
}

func (s *Stack) Peek() rune {
	return s.values[s.Length()-1]
}

func (s *Stack) Pop() (rune, error) {
	if s.Length() == 0 {
		return 0, errors.New("stack is empty")
	}
	removed := s.values[s.Length()-1]
	s.values = s.values[:s.Length()-1]
	return removed, nil
}

func makeGood(s string) string {
	var st *Stack
	st = &Stack{}
	for _, char := range s {
		if st.Length() >= 1 && st.Peek() != rune(char) && unicode.ToLower(st.Peek()) == unicode.ToLower(rune(char)) {
			st.Pop()
		} else {
			st.Push(rune(char))
		}
	}
	return string(st.values)
}

func main() {
	//testString := "leEeetcode"
	testString := "abBAcC"
	fmt.Println(makeGood(testString))
}
