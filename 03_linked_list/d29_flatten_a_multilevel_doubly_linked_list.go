package main

import "fmt"

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	var stack []*Node
	current := root
	for current != nil {
		if current.Child != nil {
			if current.Next != nil {
				stack = append(stack, current.Next)
			}
			current.Next = current.Child //
			if current.Next != nil {
				current.Next.Prev = current
			}
			current.Child = nil
		}
		if current.Next == nil && len(stack) > 0 {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			current.Next = temp
			temp.Prev = current
		}
		current = current.Next
	}
	return root
}

func main() {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n5 := &Node{Val: 5}
	n6 := &Node{Val: 6}
	n7 := &Node{Val: 7}
	n8 := &Node{Val: 8}
	n9 := &Node{Val: 9}
	n10 := &Node{Val: 10}
	n11 := &Node{Val: 11}
	n12 := &Node{Val: 12}

	n1.Next = n2
	n2.Prev = n1
	n2.Next = n3
	n3.Prev = n2
	n3.Next = n4
	n4.Prev = n3
	n4.Next = n5
	n5.Prev = n4
	n5.Next = n6
	n6.Prev = n5

	n7.Next = n8
	n8.Prev = n7
	n8.Next = n9
	n9.Prev = n8
	n9.Next = n10
	n10.Prev = n9

	n11.Next = n12
	n12.Prev = n11

	n3.Child = n7
	n8.Child = n11

	flatHead := flatten(n1)
	for curr := flatHead; curr != nil; curr = curr.Next {
		fmt.Print(curr.Val, " ")
	}
}
