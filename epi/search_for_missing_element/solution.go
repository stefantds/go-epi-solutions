package search_for_missing_element

func FindDuplicateMissing(a []int) (duplicate, missing int) {
	elemsXOR := 0
	aElemsXOR := 0
	for i, v := range a {
		elemsXOR ^= i
		aElemsXOR ^= v
	}

	// XOR-ing the elems of A and elems in [0, n] gives (missing XOR duplicate)
	missXorDup := elemsXOR ^ aElemsXOR

	bitMask := missXorDup & ^(missXorDup - 1)
	missOrDup := 0
	for i, v := range a {
		if i&bitMask > 0 {
			missOrDup ^= i
		}
		if v&bitMask > 0 {
			missOrDup ^= v
		}
	}

	for _, v := range a {
		if v == missOrDup {
			// missOrDup is the duplicated value
			return missOrDup, missXorDup ^ missOrDup
		}
	}

	// missOrDup must be the missing value
	return missXorDup ^ missOrDup, missOrDup
}
