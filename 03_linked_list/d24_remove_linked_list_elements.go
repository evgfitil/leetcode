package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var prev *ListNode
	current := head
	for current != nil && head != nil {
		if current.Val == val {
			if prev == nil {
				head = head.Next
				current = head
			} else {
				prev.Next = current.Next
				current = current.Next
			}
		} else {
			prev = current
			current = current.Next
		}
	}
	return head
}

func main() {
	//node7 := &ListNode{6, nil}
	//node6 := &ListNode{5, node7}
	//node5 := &ListNode{4, node6}
	//node4 := &ListNode{3, node5}
	//node3 := &ListNode{6, node4}
	//node2 := &ListNode{2, node3}
	//node1 := &ListNode{1, node2}
	//removeElements(node1, 6)

	//node4 := &ListNode{7, nil}
	//node3 := &ListNode{7, node4}
	//node2 := &ListNode{7, node3}
	//node1 := &ListNode{7, node2}
	//removeElements(node1, 7)

	node2 := &ListNode{2, nil}
	node1 := &ListNode{1, node2}
	removeElements(node1, 1)
}
