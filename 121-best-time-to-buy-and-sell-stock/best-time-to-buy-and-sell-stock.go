func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		newProfit := prices[i] - minPrice
		if newProfit > profit {
			profit = newProfit
		}
	}

	return profit
}