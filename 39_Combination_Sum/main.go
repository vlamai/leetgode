package _9_Combination_Sum

import (
	"fmt"
)

// Given an array of distinct integers candidates and a target integer target,
// return a list of all unique combinations of candidates where the chosen numbers sum to target.
// You may return the combinations in any order.
//
// The same number may be chosen from candidates an unlimited number of times.
// Two combinations are unique if the frequency of at least one of the chosen numbers is different.
//
// It is guaranteed that the number of unique combinations that sum up
// to target is less than 150 combinations for the given input.

func combinationSum(candidates []int, target int) [][]int {
	lastIndex := target + 1
	var result = make([][][]int, lastIndex)
	result[0] = [][]int{{}}
	for i := 0; i < lastIndex; i++ {
		current := result[i]
		if len(current) > 0 {
			for _, candidate := range candidates {
				n := i + candidate
				if n >= lastIndex {
					continue
				}
				for _, ints := range current {
					ints = append(ints, candidate)
					result[n] = append(result[n], ints)
				}
			}
		}
	}
	for i, i2 := range result {
		fmt.Printf("pos:%d c:%v\n", i, i2)
	}
	if result[target] == nil {
		return [][]int{{}}
	}
	return result[target]
}
