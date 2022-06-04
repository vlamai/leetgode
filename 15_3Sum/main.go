package _5_3Sum

import (
	"sort"
)

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
// such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.
func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1
		for l < r {
			ts := num + nums[l] + nums[r]
			if ts > 0 {
				r--
			} else if ts < 0 {
				l++
			} else {
				res = append(res, []int{num, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}

	return res
}
