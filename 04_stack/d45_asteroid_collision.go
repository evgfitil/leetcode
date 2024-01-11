// https://leetcode.com/problems/asteroid-collision/submissions/1143449602/

package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	var stack []int
	for _, asteroid := range asteroids {
		crash := false
		for len(stack) > 0 && asteroid < 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			if -asteroid > top {
				stack = stack[:len(stack)-1]
			} else if -asteroid == top {
				stack = stack[:len(stack)-1]
				crash = true
				break
			} else {
				crash = true
				break
			}
		}
		if !crash {
			stack = append(stack, asteroid)
		}
	}
	return stack
}

func main() {
	//asteroids := []int{5, 10, -5}
	asteroids := []int{10, 2, -5}
	//asteroids := []int{8, -8}
	fmt.Println(asteroidCollision(asteroids))
}
