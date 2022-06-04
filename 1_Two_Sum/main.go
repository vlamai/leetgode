package __Two_Sum

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.

func twoSum(nums []int, target int) []int {
	set := make(map[int]int)
	for i, num := range nums {
		if v, ok := set[num]; ok {
			return []int{v, i}
		} else {
			set[target-num] = i
		}
	}
	return []int{}
}
