// https://leetcode.com/problems/destination-city/description/

package main

import "fmt"

func destCity(paths [][]string) string {
	sources := make(map[string]bool)
	for _, v := range paths {
		sources[v[0]] = true
	}
	for _, v := range paths {
		if _, ok := sources[v[1]]; !ok {
			return v[1]
		}
	}
	return ""
}

func main() {
	testPaths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	fmt.Println(destCity(testPaths))
}
