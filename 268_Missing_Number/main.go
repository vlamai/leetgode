package _68_Missing_Number

//
// Given an array nums containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.
//
func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, num := range nums {
		sum -= num
	}
	return sum
}
