// https://leetcode.com/problems/design-linked-list/description/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	Head *ListNode
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{Head: nil, Size: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index > this.Size {
		return -1
	}
	current := this.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.Head = &ListNode{Val: val, Next: this.Head}
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{Val: val, Next: nil}
	if this.Size == 0 {
		this.Head = newNode
	} else {
		current := this.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}

	newNode := &ListNode{Val: val, Next: nil}
	current := this.Head
	if index == 0 {
		newNode.Next = current
		this.Head = newNode
	} else {
		var prev *ListNode
		for i := 0; i < index; i++ {
			prev = current
			current = current.Next
		}
		prev.Next = newNode
		newNode.Next = current
	}
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	switch {
	case index >= this.Size:
		return
	case index == 0:
		this.Head = this.Head.Next
		this.Size--
	default:
		var prev *ListNode
		current := this.Head
		for i := 0; i < index; i++ {
			prev = current
			current = current.Next
		}
		prev.Next = current.Next
		this.Size--
	}
}

func main() {
	obj := Constructor()
	obj.AddAtHead(7)
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	obj.AddAtIndex(3, 0)
	obj.DeleteAtIndex(2)
	obj.AddAtHead(6)
	obj.AddAtTail(4)
	obj.Get(4)
	obj.AddAtHead(4)
	obj.AddAtIndex(5, 0)
	obj.AddAtHead(6)
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
