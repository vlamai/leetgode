package _38_Product_of_Array_Except_Self

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements
// of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// You must write an algorithm that runs in O(n) time and without using the division operation.

func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
	lastIndex := numsLen - 1
	l := make([]int, numsLen)
	r := make([]int, numsLen)
	for i := 0; i < numsLen; i++ {
		l[i] = nums[i]
		if i > 0 {
			l[i] = nums[i] * l[i-1]
		}
	}

	for i := lastIndex; i >= 0; i-- {
		r[i] = nums[i]
		if i != lastIndex {
			r[i] = nums[i] * r[i+1]
		}
	}
	var res = make([]int, numsLen)
	for i := 0; i < numsLen; i++ {
		lv := 1
		rv := 1
		if i != 0 {
			lv = l[i-1]
		}
		if i != lastIndex {
			rv = r[i+1]
		}
		res[i] = lv * rv
	}
	return res
}
