// https://leetcode.com/problems/dota2-senate/

package main

import (
	"container/list"
	"fmt"
)

func predictPartyVictory(senate string) string {
	radiant := list.New()
	dire := list.New()
	senatorsCount := len(senate)
	for idx, senator := range senate {
		if string(senator) == "R" {
			radiant.PushBack(idx)
		} else {
			dire.PushBack(idx)
		}
	}
	for radiant.Len() > 0 && dire.Len() > 0 {
		senR := radiant.Front().Value.(int)
		senD := dire.Front().Value.(int)
		if senR < senD {
			newRadiant := senR + senatorsCount
			radiant.PushBack(newRadiant)
		} else {
			newDire := senD + senatorsCount
			dire.PushBack(newDire)
		}
		dire.Remove(dire.Front())
		radiant.Remove(radiant.Front())
	}
	if radiant.Len() > dire.Len() {
		return "Radiant"
	} else {
		return "Dire"
	}
}

func main() {
	//testSenate := "RDD" // -> D
	//testSenate := "RD" // -> R
	//testSenate := "RRDDD" // -> D
	testSenate := "DDRRR" // -> Dire
	fmt.Println(predictPartyVictory(testSenate))
}
