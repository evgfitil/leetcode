// https://leetcode.com/problems/reverse-linked-list-ii/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var prev, third *ListNode
	current := head
	for left > 1 {
		prev = current
		current = current.Next
		left--
		right--
	}
	tail, con := current, prev
	for right > 0 {
		third = current.Next
		current.Next = prev
		prev = current
		current = third
		right--
	}

	if con != nil {
		con.Next = prev
	} else {
		head = prev
	}
	tail.Next = current
	return head
}
func main() {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	reverseBetween(node1, 2, 4)
}
