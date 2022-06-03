package _48_Find_All_Numbers_Disappeared_in_an_Array

//
// Given an array nums of n integers where nums[i] is in the range [1, n],
// return an array of all the integers in the range [1, n] that do not appear in nums.
//
func findDisappearedNumbers(nums []int) []int {
	var result []int
	for _, num := range nums {
		if num < 0 {
			num *= -1
		}
		if nums[num-1] < 0 {
			continue
		}
		nums[num-1] *= -1
	}
	for i, num := range nums {
		if num > 0 {
			result = append(result, i+1)
		}
	}
	return result
}
