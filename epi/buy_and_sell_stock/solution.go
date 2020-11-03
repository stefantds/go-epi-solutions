package buy_and_sell_stock

func ComputeMaxProfit(prices []float64) float64 {
	if len(prices) == 0 {
		return 0.0
	}
	minSoFar := prices[0]
	maxProfit := 0.0

	for i := 1; i < len(prices); i++ {
		profit := prices[i] - minSoFar
		if profit > maxProfit {
			maxProfit = profit
		}
		if prices[i] < minSoFar {
			minSoFar = prices[i]
		}
	}

	return maxProfit
}
