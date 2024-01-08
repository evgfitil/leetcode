// https://leetcode.com/problems/min-stack/description/

package main

import "fmt"

type stackItem struct {
	value    int
	minValue int
}

type MinStack struct {
	items []stackItem
}

func Constructor() MinStack {
	return MinStack{items: make([]stackItem, 0)}
}

func (this *MinStack) Push(val int) {
	minVal := val
	if len(this.items) > 0 && val > this.GetMin() {
		minVal = this.GetMin()
	}
	this.items = append(this.items, stackItem{value: val, minValue: minVal})
}

func (this *MinStack) Pop() {
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1].value
}

func (this *MinStack) GetMin() int {
	return this.items[len(this.items)-1].minValue
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	//obj.Push(-2)
	//obj.Push(0)
	//obj.Push(-3)
	//fmt.Println(obj.GetMin())
	//obj.Pop()
	//fmt.Println(obj.Top())
	//fmt.Println(obj.GetMin())
	//
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-1)
	fmt.Println(obj.GetMin())
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.GetMin())
}
