package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	current := node
	var prev *ListNode
	for current.Next != nil {
		prev = current
		current.Val = current.Next.Val
		current = current.Next
	}
	prev.Next = nil
}

func main() {
	node4 := &ListNode{9, nil}
	node3 := &ListNode{1, node4}
	node2 := &ListNode{5, node3}
	node1 := &ListNode{4, node2}
	fmt.Println(node1)
	deleteNode(node2)
}
