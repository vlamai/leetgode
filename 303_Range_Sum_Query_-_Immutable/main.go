package _03_Range_Sum_Query___Immutable

// Given an integer array nums, handle multiple queries of the following type:
//
//    1. Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
//
// Implement the NumArray class:
//
//    1. NumArray(int[] nums) Initializes the object with the integer array nums.
//    2. int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right
//    inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

type NumArray struct {
}

func Constructor(nums []int) NumArray {
	return NumArray{}
}

func (this *NumArray) SumRange(left int, right int) int {
	return -1
}
