// https://leetcode.com/problems/online-stock-span/description/

package main

import "fmt"

type stock struct {
	price int
	span  int
}
type StockSpanner struct {
	stack []stock
}

func Constructor() StockSpanner {
	return StockSpanner{stack: make([]stock, 0)}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.stack) > 0 && price >= this.stack[len(this.stack)-1].price {
		span += this.stack[len(this.stack)-1].span
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, stock{price: price, span: span})
	return this.stack[len(this.stack)-1].span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

func main() {
	obj := Constructor()
	fmt.Println(obj.Next(100))
	fmt.Println(obj.Next(80))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(70))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(75))
	fmt.Println(obj.Next(85))
}
