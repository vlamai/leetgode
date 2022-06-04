package _52_Maximum_Product_Subarray

// Given an integer array nums, find a contiguous non-empty subarray within the array that has the largest product,
// and return the product.
//
// The test cases are generated so that the answer will fit in a 32-bit integer.
//
// A subarray is a contiguous subsequence of the array.

func maxProduct(nums []int) int {
	maxProd := nums[0]
	curMin, curMax := 1, 1
	for _, num := range nums {
		if num == 0 {
			curMin, curMax = 1, 1
		}
		tmpMax := num * curMax
		tmpMin := num * curMin
		curMax = max(tmpMax, tmpMin, num)
		curMin = min(tmpMax, tmpMin, num)
		maxProd = max(curMax, maxProd)
	}
	return maxProd
}

func min(i ...int) int {
	m := i[0]
	for _, n := range i {
		if m > n {
			m = n
		}
	}
	return m
}

func max(i ...int) int {
	m := i[0]
	for _, n := range i[1:] {
		if n > m {
			m = n
		}
	}
	return m
}
