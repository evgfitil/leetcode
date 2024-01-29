package main

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func reverseList(head *listNode) *listNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversedHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return reversedHead
}

func main() {
	node5 := &listNode{5, nil}
	node4 := &listNode{4, node5}
	node3 := &listNode{3, node4}
	node2 := &listNode{2, node3}
	node1 := &listNode{1, node2}
	fmt.Println(reverseList(node1))
}
