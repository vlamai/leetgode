package _6_Plus_One

import "fmt"

func plusOne(digits []int) []int {
	var res = make([]int, 0, len(digits)+1)
	var n, c = addOne(digits[len(digits)-1], 1)
	fmt.Println(n, c)
	i := len(digits) - 2
	for c != 0 && i >= 0 {
		res = prependInt(res, c)
		currentDigit := digits[i]
		n, c = addOne(currentDigit, n)
		fmt.Println(n, c)
		i--
	}
	if c != 0 {
		res = prependInt(res, c)
	}
	return res
}

func addOne(n, a int) (int, int) {
	n += a
	return n / 10, n % 10
}
func prependInt(x []int, y int) []int {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}
