package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeSize := 0
	left := head
	for left != nil {
		nodeSize++
		left = left.Next
	}
	if nodeSize == 1 && n == 1 {
		return nil
	}
	index := nodeSize - n
	if index == 0 {
		head = head.Next
	} else {
		var prev *ListNode
		current := head
		for i := 0; i < index; i++ {
			prev = current
			current = current.Next
		}
		prev.Next = current.Next
	}
	return head
}

func main() {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	removeNthFromEnd(node1, 2)
}
