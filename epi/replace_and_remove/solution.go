package replace_and_remove

func ReplaceAndRemove(size int, s []rune) int {
	aCount, writeIdx := 0, 0

	// remove 'b's and count 'a's
	for i := 0; i < size; i++ {
		if s[i] == 'a' {
			aCount++
		}
		if s[i] != 'b' {
			s[writeIdx] = s[i]
			writeIdx++
		}
	}

	// replace 'a's by two 'd's starting from the end
	readIdx := writeIdx - 1
	writeIdx = writeIdx + aCount - 1
	finalSize := writeIdx + 1

	for readIdx >= 0 {
		if s[readIdx] == 'a' {
			s[writeIdx] = 'd'
			s[writeIdx-1] = 'd'
			writeIdx -= 2
		} else {
			s[writeIdx] = s[readIdx]
			writeIdx--
		}
		readIdx--
	}

	return finalSize
}
