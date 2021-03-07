package making_change

var Coins = []int{
	100, 50, 25, 10, 5, 1,
}

func ChangeMaking(cents int) int {
	nbCoins := 0
	for _, currentValue := range Coins {
		nbCoins += cents / currentValue
		cents = cents % currentValue
	}

	return nbCoins
}
