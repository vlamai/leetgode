package _53_Find_Minimum_in_Rotated_Sorted_Array

// Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
// For example, the array nums = [0,1,2,4,5,6,7] might become:
//
//    [4,5,6,7,0,1,2] if it was rotated 4 times.
//    [0,1,2,4,5,6,7] if it was rotated 7 times.
//
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
//
// Given the sorted rotated array nums of unique elements, return the minimum element of this array.
//
// You must write an algorithm that runs in O(log n) time.

func findMin(nums []int) int {
	min := nums[0]
	l := 0
	r := len(nums) - 1
	for l <= r {
		if nums[l] < nums[r] {
			if min > nums[l] {
				min = nums[l]
			}
			break
		}
		mid := (l + r) / 2
		if nums[mid] >= nums[l] {
			l = mid + 1
		} else {
			r = mid - 1
		}
		if min > nums[mid] {
			min = nums[mid]
		}
	}

	return min
}
