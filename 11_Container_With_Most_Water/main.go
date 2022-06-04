package _1_Container_With_Most_Water

// You are given an integer array height of length n.
// There are n vertical lines drawn such that the two endpoints
// of the ith line are (i, 0) and (i, height[i]).
//
// Find two lines that together with the x-axis form a container,
// such that the container contains the most water.
//
// Return the maximum amount of water a container can store.
//
// Notice that you may not slant the container.
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxA := 0
	for l <= r {
		leftHeight := height[l]
		rightHeight := height[r]
		area := (r - l) * min(leftHeight, rightHeight)
		maxA = max(area, maxA)
		if leftHeight >= rightHeight {
			r--
		} else {
			l++
		}
	}
	return maxA
}

func min(nums ...int) int {
	m := nums[0]
	for _, num := range nums {
		if m > num {
			m = num
		}
	}
	return m
}

func max(nums ...int) int {
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}
