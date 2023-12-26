package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	lengthNode := 1
	var left *ListNode
	current := head
	for current != nil {
		if lengthNode == k {
			left = current
		}
		current = current.Next
		lengthNode++
	}
	right := head
	rightPoint := lengthNode - k
	for rightPoint != 1 {
		right = right.Next
		rightPoint--
	}
	left.Val, right.Val = right.Val, left.Val
	return head
}

func printLinkedList(head *ListNode) {
	var values []string
	for head != nil {
		values = append(values, strconv.Itoa(head.Val))
		head = head.Next
	}
	fmt.Println(strings.Join(values, ", "))
}

func main() {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	printLinkedList(swapNodes(node1, 2))
}
