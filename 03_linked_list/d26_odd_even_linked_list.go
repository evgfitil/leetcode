package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	left, right := head, head.Next
	newHead := right
	for right != nil && right.Next != nil {
		left.Next = right.Next
		left = left.Next
		right.Next = left.Next
		right = right.Next
	}
	left.Next = newHead
	return head
}

func main() {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	oddEvenList(node1)
}

// 1 2 3 4 5
// o e
//   e o
