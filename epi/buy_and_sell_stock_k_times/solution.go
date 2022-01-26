package buy_and_sell_stock_k_times

import "math"

func BuyAndSellStockKTimes(prices []float64, k int) float64 {
	if k == 0 {
		return 0
	}
	if 2*k >= len(prices) {
		return unlimitedTransactionsProfit(prices)
	}

	minPrices := make([]float64, k)
	maxProfits := make([]float64, k)
	for i := range minPrices {
		minPrices[i] = math.MaxFloat64
	}

	for _, p := range prices {
		for i := k - 1; i >= 0; i-- {
			if p-minPrices[i] > maxProfits[i] {
				maxProfits[i] = p - minPrices[i]
			}
			var price float64
			if i == 0 {
				price = p
			} else {
				price = p - maxProfits[i-1]
			}
			if price < minPrices[i] {
				minPrices[i] = price
			}
		}
	}

	return maxProfits[k-1]
}

func unlimitedTransactionsProfit(prices []float64) float64 {
	profit := 0.0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
