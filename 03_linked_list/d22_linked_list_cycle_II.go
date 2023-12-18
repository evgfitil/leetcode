package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	indexMap := make(map[*ListNode]bool)
	left := head
	for left != nil && left.Next != nil {
		if _, ok := indexMap[left]; !ok {
			indexMap[left] = true
		} else {
			return left
		}
		left = left.Next
	}
	return nil
}

func main() {
	head := &ListNode{Val: 1}
	second := &ListNode{Val: 2}

	head.Next = second
	second.Next = head
	loopStart := detectCycle(head)
	fmt.Println(loopStart)
}
