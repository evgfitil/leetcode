// https://leetcode.com/problems/palindrome-linked-list/description/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	left := head
	right := reverseList(middleNode(head))
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left, right = left.Next, right.Next
	}
	return true
}

func middleNode(head *ListNode) *ListNode {
	left, right := head, head
	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next.Next
	}
	return left
}

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

func main() {
	node4 := &ListNode{1, nil}
	node3 := &ListNode{2, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	fmt.Println(isPalindrome(node1))
}

// 1 2 2 1
//
