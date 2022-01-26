package rook_attack

func RookAttack(a [][]int) {
	var firstColumnHasZero, firstRowHasZero bool
	for i := 0; i < len(a); i++ {
		if a[i][0] == 0 {
			firstColumnHasZero = true
			break
		}
	}
	for i := 0; i < len(a[0]); i++ {
		if a[0][i] == 0 {
			firstRowHasZero = true
			break
		}
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] == 0 {
				// set the corresponding position in the
				// first row and column to 0
				a[i][0] = 0
				a[0][j] = 0
			}
		}
	}

	for i := 1; i < len(a); i++ {
		for j := 1; j < len(a[i]); j++ {
			if a[0][j] == 0 || a[i][0] == 0 {
				a[i][j] = 0
			}
		}
	}

	if firstColumnHasZero {
		for i := 0; i < len(a); i++ {
			a[i][0] = 0
		}
	}

	if firstRowHasZero {
		for i := 0; i < len(a[0]); i++ {
			a[0][i] = 0
		}
	}
}
