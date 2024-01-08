// https://leetcode.com/problems/sliding-window-maximum/description/

package main

import "fmt"

type Deque struct {
	items []int
}

func (d *Deque) PushFront(item int) {
	d.items = append([]int{item}, d.items...)
}

func (d *Deque) PushBack(item int) {
	d.items = append(d.items, item)
}

func (d *Deque) PopFront() (int, bool) {
	if len(d.items) == 0 {
		return 0, false
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item, true
}

func (d *Deque) PopBack() (int, bool) {
	if len(d.items) == 0 {
		return 0, false
	}
	index := len(d.items) - 1
	item := d.items[index]
	d.items = d.items[:index]
	return item, true
}

func (d *Deque) Front() int {
	return d.items[0]
}

func (d *Deque) Back() int {
	return d.items[len(d.items)-1]
}

func maxSlidingWindow(nums []int, k int) []int {
	dq := &Deque{}
	var res []int
	for i, num := range nums {
		if len(dq.items) > 0 && dq.Front() == i-k {
			dq.PopFront()
		}
		for len(dq.items) > 0 && nums[dq.Back()] < num {
			dq.PopBack()
		}
		dq.PushBack(i)
		if i >= k-1 {
			res = append(res, nums[dq.Front()])
		}
	}
	return res
}

func main() {
	testNums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	//testNums := []int{7, 7, 3, 4}
	//k := 3
	fmt.Println(maxSlidingWindow(testNums, k))
}

// 7 2 4
// 2
// 7 2
// 2 4

// 7 7 3 4
// 3
// 7 7 3
// 7 3 4
