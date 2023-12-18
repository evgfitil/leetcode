// https://leetcode.com/problems/permutation-in-string/description/

package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	s1Map := make(map[rune]int)
	s2Map := make(map[rune]int)

	for _, value := range s1 {
		s1Map[value] += 1
	}

	left, windowSize := 0, len(s1)
	for right := 0; right < len(s2); right++ {
		currentChar := s2[right]
		if _, ok := s1Map[rune(currentChar)]; ok {
			s2Map[rune(currentChar)] += 1
		}
		if right-left+1 == windowSize {
			if compareMaps(s1Map, s2Map) {
				return true
			}
			leftChar := s2[left]
			if _, ok := s2Map[rune(leftChar)]; ok {
				s2Map[rune(leftChar)] -= 1
			}
			left++
		}
	}
	return false
}

func compareMaps(map1, map2 map[rune]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, value1 := range map1 {
		if value2, ok := map2[key]; !ok || value1 != value2 {
			return false
		}
	}
	return true
}

func main() {
	testS1, testS2 := "ab", "eidbaooo"
	fmt.Println(checkInclusion(testS1, testS2))
}
