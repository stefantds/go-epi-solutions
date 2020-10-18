package power_set

func GeneratePowerSet(inputSet []int) [][]int {
	return generatePowerSetHelper(inputSet, 0)
}

func generatePowerSetHelper(inputSet []int, pos int) [][]int {
	if pos == len(inputSet) {
		return [][]int{{}}
	}
	currentVal := inputSet[pos]
	ps := generatePowerSetHelper(inputSet, pos+1)

	result := make([][]int, 0)
	for _, v := range ps {
		vCopy := make([]int, len(v))
		copy(vCopy, v)
		result = append(result, vCopy)                 // sets without the current element
		result = append(result, append(v, currentVal)) // sets with the current element
	}
	return result
}
