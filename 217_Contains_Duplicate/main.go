package _17_Contains_Duplicate

//
// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.
//
func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
