package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func swapPairs(head *listNode) *listNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := head.Next
	first.Next = swapPairs(second.Next)
	second.Next = first
	return second
}

func main() {
	node4 := &listNode{4, nil}
	node3 := &listNode{3, node4}
	node2 := &listNode{2, node3}
	node1 := &listNode{1, node2}
	fmt.Println(swapPairs(node1))
}
