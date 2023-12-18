// https://leetcode.com/problems/linked-list-cycle/description/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	left, right := head, head
	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next.Next
		if left == right {
			return true
		}
	}
	return false
}

func main() {
	head := &ListNode{Val: 1}
	second := &ListNode{Val: 2}
	third := &ListNode{Val: 3}
	fourth := &ListNode{Val: 4}

	head.Next = second
	second.Next = third
	third.Next = fourth

	fourth.Next = second
	fmt.Println(hasCycle(head))
}
