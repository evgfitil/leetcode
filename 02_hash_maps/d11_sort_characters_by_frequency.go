// https://leetcode.com/problems/sort-characters-by-frequency/description/

package main

import (
	"fmt"
	"sort"
	"strings"
)

type sliceCount struct {
	Key   rune
	Value int
}

func frequencySort(s string) string {
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}
	var sc []sliceCount
	for key, value := range counts {
		sc = append(sc, sliceCount{key, value})
	}
	sort.Slice(sc, func(i, j int) bool {
		return sc[i].Value > sc[j].Value
	})
	var res strings.Builder
	for _, kv := range sc {
		res.WriteString(strings.Repeat(string(kv.Key), kv.Value))
	}
	return res.String()
}

func main() {
	testString := "tree"
	fmt.Println(frequencySort(testString))
}
