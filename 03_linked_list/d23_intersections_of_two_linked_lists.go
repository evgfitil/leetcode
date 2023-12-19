// https://leetcode.com/problems/intersection-of-two-linked-lists/description/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//// Hash map approach, time: O(m + n), memory: O(n)
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	if headA == headB {
//		return headA
//	}
//	nodesMap := make(map[*ListNode]bool)
//	first := headA
//	for first != nil {
//		nodesMap[first] = true
//		first = first.Next
//	}
//	second := headB
//	for second != nil {
//		if _, ok := nodesMap[second]; ok {
//			return second
//		}
//		second = second.Next
//	}
//	return nil
//}

// Two pointers approach, time: O(m+n), memory: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	left, right := headA, headB
	for left != right {
		if left == nil {
			left = headA
		}
		if right == nil {
			right = headB
		}
		if left == right {
			return left
		}
		left = left.Next
		right = right.Next
	}
	return left
}

func main() {
	nodeC3 := &ListNode{5, nil}
	nodeC2 := &ListNode{4, nodeC3}
	nodeC1 := &ListNode{8, nodeC2}

	nodeA1 := &ListNode{1, nodeC1}
	headA := &ListNode{4, nodeA1}

	nodeB2 := &ListNode{1, nodeC1}
	nodeB1 := &ListNode{6, nodeB2}
	headB := &ListNode{5, nodeB1}

	fmt.Println(getIntersectionNode(headA, headB))
}
