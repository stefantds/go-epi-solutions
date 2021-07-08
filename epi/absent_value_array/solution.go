package absent_value_array

// ResetIterator has a metgod Iterator that returns a new channel
// containing the same stream of values every time.
// It effectively allows to reset the channel and read again from it multiple times.
type ResetIterator interface {
	Iterator() <-chan int32
}

func FindMissingElement(stream ResetIterator) int32 {
	counter := make([]int32, 1<<16)

	// count how many each first half of the word appears in the stream
	values := stream.Iterator()
	for v := range values {
		idx := v >> 16
		counter[idx]++
	}

	// find a value that appeared less than pow(2, 16) times in the stream
	var candidatePrefix int32
	for i, v := range counter {
		if v < (1 << 16) {
			candidatePrefix = int32(i)
			break
		}
	}

	// get all the suffixes of values starting with the candidatePrefix
	presentValues := make([]bool, 1<<16)
	values = stream.Iterator()
	for v := range values {
		prefix := v >> 16
		if prefix == candidatePrefix {
			suffix := (v << 16) >> 16
			presentValues[suffix] = true
		}
	}

	// find the first suffix not in the seen values
	for i := int32(0); i < 1<<16; i++ {
		if !presentValues[i] {
			return candidatePrefix<<16 | i
		}
	}

	panic("unreachable")
}
