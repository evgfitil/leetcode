// https://leetcode.com/problems/implement-stack-using-queues/description/

package main

import "fmt"

type MyStack struct {
	queue1 []int
	queue2 []int
	top    int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue1 = append(this.queue1, x)
	this.top = x
}

func (this *MyStack) Pop() int {
	for len(this.queue1) > 1 {
		temp := this.queue1[0]
		this.queue1 = this.queue1[1:]
		this.queue2 = append(this.queue2, temp)
	}
	popped := this.queue1[0]
	this.queue1 = []int{}
	for len(this.queue2) > 0 {
		temp := this.queue2[0]
		this.queue1 = append(this.queue1, temp)
		this.queue2 = this.queue2[1:]
		this.top = temp
	}
	return popped
}

func (this *MyStack) Top() int {
	return this.top
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}
