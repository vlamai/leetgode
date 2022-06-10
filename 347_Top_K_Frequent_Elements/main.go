package _47_Top_K_Frequent_Elements

func topKFrequent(nums []int, k int) []int {
	var m = make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	var c = make([][]int, len(nums)+1)
	for num, count := range m {
		c[count] = append(c[count], num)
	}

	var res = make([]int, 0, k)
	for i := len(c) - 1; i > 0; i-- {
		if len(c[i]) > 0 {
			res = append(res, c[i]...)
		}
	}

	return res[:k]
}
