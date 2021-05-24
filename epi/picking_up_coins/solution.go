package picking_up_coins

import (
	utils "github.com/stefantds/go-epi-judge/test_utils"
)

func PickUpCoins(coins []int) int {
	cache := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		cache[i] = make([]int, len(coins))
	}
	return chooseCoins(coins, 0, len(coins)-1, cache)
}

func chooseCoins(coins []int, a, b int, cache [][]int) int {
	if a > b {
		return 0
	}

	if cache[a][b] == 0 {
		valueA := coins[a] + utils.Min(chooseCoins(coins, a+2, b, cache), chooseCoins(coins, a+1, b-1, cache))
		valueB := coins[b] + utils.Min(chooseCoins(coins, a, b-2, cache), chooseCoins(coins, a+1, b-1, cache))

		cache[a][b] = utils.Max(valueA, valueB)
	}

	return cache[a][b]
}
