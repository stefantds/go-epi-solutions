package buy_and_sell_stock_twice

import (
	"math"
)

func BuyAndSellStockTwice(prices []float64) float64 {
	firstBuyProfits := make([]float64, len(prices))
	minPriceSoFar := math.MaxFloat64
	maxProfit := 0.0

	for i, p := range prices {
		minPriceSoFar = math.Min(p, minPriceSoFar)
		maxProfit = math.Max(maxProfit, p-minPriceSoFar)
		firstBuyProfits[i] = maxProfit
	}

	maxPriceSoFar := -math.MaxFloat64
	for i := len(prices) - 1; i > 0; i-- {
		maxPriceSoFar = math.Max(prices[i], maxPriceSoFar)
		profit := firstBuyProfits[i-1] /*profit from first buy/sell*/ + maxPriceSoFar - prices[i] /* profit from second buy/sell*/
		maxProfit = math.Max(profit, maxProfit)
	}

	return maxProfit
}
