package _21_Best_Time_to_Buy_and_Sell_Stock

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock and choosing
// a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// 7, 1, 5, 3, 6, 4
func maxProfit(prices []int) int {
	minPrice := prices[0]
	var maxProfit = 0
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		}
		curProfit := price - minPrice
		if curProfit > maxProfit {
			maxProfit = curProfit
		}
	}
	return maxProfit
}
