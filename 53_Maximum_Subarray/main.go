package _3_Maximum_Subarray

// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
//
// A subarray is a contiguous part of an array.

// -2, 1, -3, 4, -1, 2, 1, -5, 4
//            4, -1, 2, 1
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	cSum := 0
	for _, num := range nums {
		if cSum < 0 {
			cSum = 0
		}
		cSum += num
		if maxSum < cSum {
			maxSum = cSum
		}
	}
	return maxSum
}
