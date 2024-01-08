// https://leetcode.com/problems/number-of-recent-calls/description/

package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	First *Node
	Last  *Node
	Len   int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Size() int {
	return q.Len
}

func (q *Queue) Put(value int) {
	newNode := &Node{Value: value}

	if q.Last == nil {
		q.First = newNode
		q.Last = newNode
	} else {
		q.Last.Next = newNode
		q.Last = newNode
	}
	q.Len++
}

func (q *Queue) Get() (int, error) {
	if q.First == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	temp := q.First
	q.First = temp.Next
	q.Len--
	if q.Len == 0 {
		q.Last = nil
	}
	return temp.Value, nil
}

type RecentCounter struct {
	queue *Queue
}

func Constructor() RecentCounter {
	return RecentCounter{queue: NewQueue()}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue.Put(t)
	for this.queue.First != nil && this.queue.First.Value < t-3000 {
		this.queue.Get()
	}
	return this.queue.Size()
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))    // 1
	fmt.Println(obj.Ping(100))  // 2
	fmt.Println(obj.Ping(3001)) // 3
	fmt.Println(obj.Ping(3002)) // 3
}
