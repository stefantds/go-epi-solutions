package number_of_score_combinations

func NumCombinationsForFinalScore(finalScore int, individualPlayScores []int) int {
	numCombinationsForScore := make([][]int, len(individualPlayScores))
	for i := 0; i < len(individualPlayScores); i++ {
		numCombinationsForScore[i] = make([]int, finalScore+1)
	}

	for i := 0; i < len(individualPlayScores); i++ {
		numCombinationsForScore[i][0] = 1
		for j := 1; j <= finalScore; j++ {
			withoutThisPlay, withThisPlay := 0, 0
			if i-1 >= 0 {
				withoutThisPlay = numCombinationsForScore[i-1][j]
			}
			if j-individualPlayScores[i] >= 0 {
				withThisPlay = numCombinationsForScore[i][j-individualPlayScores[i]]
			}
			numCombinationsForScore[i][j] = withThisPlay + withoutThisPlay
		}
	}

	return numCombinationsForScore[len(individualPlayScores)-1][finalScore]
}
