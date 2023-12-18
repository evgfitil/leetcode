// https://leetcode.com/problems/middle-of-the-linked-list/description/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//func lengthNode(head *ListNode) int {
//	length := 0
//	current := head
//	for current != nil {
//		length++
//		current = current.Next
//	}
//	return length
//}

func middleNode(head *ListNode) *ListNode {
	left, right := head, head
	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next.Next
	}
	return left
}

func main() {
	node6 := &ListNode{2, nil}
	node5 := &ListNode{1, node6}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	fmt.Println(middleNode(node1))
}

// 1 2 3 4 5 6
//   l     r
//     l r
// l = 2, r = 3; 2 + 3 = 5; 6 - 5 = 1

// 1 2 3 4 5
//   l   r
//     l == r
