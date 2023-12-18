// https://leetcode.com/problems/reverse-linked-list/description/

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

// recursive approach
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversedHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return reversedHead
}

// iterative approach
func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	current := head
	for current != nil {
		nextTemp := current.Next
		current.Next = previous
		previous = current
		current = nextTemp
	}
	return previous
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
	//printLinkedList(reverseList(node1))
	printLinkedList(reverseListRecursive(node1))
}

// 1 -> 2 -> 3 -> 4 -> 5 -> nil
// 5 -> 4 -> 3 -> 2 -> 1 -> nil
