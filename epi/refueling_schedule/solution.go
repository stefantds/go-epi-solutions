package refueling_schedule

const MPG int = 20

func FindAmpleCity(gallons []int, distances []int) int {
	currentGas := 0 // we don't care about the actual value, we just need to find the minimum
	minGas := 0
	minCity := 0
	numCities := len(distances)

	for i := 1; i < numCities; i++ {
		// currentGas is the remaining gas when arriving at city i
		currentGas = currentGas + gallons[i-1] - gasUsed(distances[i-1])
		if currentGas < minGas {
			minGas = currentGas
			minCity = i
		}
	}

	return minCity
}

func gasUsed(distance int) int {
	return distance / MPG
}
