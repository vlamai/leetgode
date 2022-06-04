package _22_Coin_Change

// You are given an integer array coins representing coins of different denominations and
// an integer amount representing a total amount of money.
//
// Return the fewest number of coins that you need to make up that amount.
// If that amount of money cannot be made up by any combination of the coins, return -1.
//
// You may assume that you have an infinite number of each kind of coin.

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for a := 1; a <= amount; a++ {
		for _, coin := range coins {
			if a-coin >= 0 {
				o := 1 + dp[a-coin]
				if dp[a] > o {
					dp[a] = o
				}
			}
		}
	}
	res := dp[amount]
	if res == amount+1 {
		return -1
	}
	return res
}
