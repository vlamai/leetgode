package _69_Majority_Element

func majorityElement(nums []int) int {
	var result int
	var vote int

	for _, num := range nums {
		if vote == 0 {
			result = num
		}
		if num == result {
			vote++
		} else {
			vote--
		}
	}
	return result
}
