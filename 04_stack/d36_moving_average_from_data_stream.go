package main

import "fmt"

//// Array approach
//type MovingAverage struct {
//	Size  int
//	Queue []int
//	Sum   int
//}
//
//func Constructor(size int) MovingAverage {
//	return MovingAverage{
//		Size:  size,
//		Queue: make([]int, 0, size),
//	}
//}
//
//func (this *MovingAverage) Next(val int) float64 {
//	this.Queue = append(this.Queue, val)
//	this.Sum += val
//
//	if len(this.Queue) > this.Size {
//		this.Sum -= this.Queue[0]
//		this.Queue = this.Queue[1:]
//	}
//
//	return float64(this.Sum) / float64(len(this.Queue))
//}
//
///**
// * Your MovingAverage object will be instantiated and called as such:
// * obj := Constructor(size);
// * param_1 := obj.Next(val);
// */

typ

func main() {
	ma := Constructor(3)
	fmt.Println(ma.Next(1))
	fmt.Println(ma.Next(10))
	fmt.Println(ma.Next(3))
	fmt.Println(ma.Next(5))
}
