package _0_Climbing_Stairs

// You are climbing a staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	f := 1
	s := 2
	for i := 2; i < n; i++ {
		c := f + s
		f = s
		s = c
	}
	return s
}
