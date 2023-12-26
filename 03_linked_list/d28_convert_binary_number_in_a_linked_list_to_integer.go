package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	res := 0
	current := head
	for current != nil {
		res = res*2 + current.Val
		current = current.Next
	}
	return res
}

func main() {
	node3 := &ListNode{1, nil}
	node2 := &ListNode{0, node3}
	node1 := &ListNode{1, node2}
	fmt.Println(getDecimalValue(node1))
}
