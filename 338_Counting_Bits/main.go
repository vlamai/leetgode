package _38_Counting_Bits

// Given an integer n, return an array ans of length n + 1
// such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

func countBits(n int) []int {
	result := []int{0}
	for i := 1; i <= n; i++ {
		result = append(result, result[i>>1]+i%2)
	}
	return result
}
