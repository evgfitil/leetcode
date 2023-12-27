package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	current := head
	var items []int
	pairSumMax := 0
	for current != nil {
		items = append(items, current.Val)
		current = current.Next
	}
	n := len(items)
	left, right := 0, n-1
	for left < right {
		currentMax := items[left] + items[right]
		pairSumMax = max(currentMax, pairSumMax)
		left++
		right--
	}
	//for idx, value := range items {
	//	currentMax := value + items[n-idx-1]
	//	pairSumMax = max(currentMax, pairSumMax)
	//}
	return pairSumMax
}

func main() {
	node4 := &ListNode{1, nil}
	node3 := &ListNode{2, node4}
	node2 := &ListNode{4, node3}
	node1 := &ListNode{5, node2}
	fmt.Println(pairSum(node1))
}
