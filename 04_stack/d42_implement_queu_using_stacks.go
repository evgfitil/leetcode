// https://leetcode.com/problems/implement-queue-using-stacks/description/

package main

import "fmt"

type MyQueue struct {
	front []int
	back  []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.front = append(this.front, x)
}

func (this *MyQueue) Pop() int {
	if len(this.back) == 0 {
		for len(this.front) > 0 {
			top := this.front[len(this.front)-1]
			this.front = this.front[:len(this.front)-1]
			this.back = append(this.back, top)
		}
	}
	currentTop := this.back[len(this.back)-1]
	this.back = this.back[:len(this.back)-1]
	return currentTop
}

func (this *MyQueue) Peek() int {
	if len(this.back) == 0 {
		for len(this.front) > 0 {
			top := this.front[len(this.front)-1]
			this.front = this.front[:len(this.front)-1]
			this.back = append(this.back, top)
		}
	}
	currentTop := this.back[len(this.back)-1]
	return currentTop
}

func (this *MyQueue) Empty() bool {
	if len(this.front) == 0 && len(this.back) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Empty())
}

// [] []
// [1] [2]
// []
