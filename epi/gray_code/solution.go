package gray_code

import (
	"fmt"
)

func GrayCode(numBits int) []int64 {
	return findGrayCode(numBits, []int64{0}, map[int64]bool{0: true})
}

func findGrayCode(numBits int, currentPath []int64, existingValues map[int64]bool) []int64 {
	lastElem := currentPath[len(currentPath)-1]

	if len(existingValues) == 1<<numBits {
		if len(existingValues) == 1 || differsByOne(lastElem, currentPath[0]) {
			return currentPath // the current path is a valid gray code
		} else {
			return nil
		}
	}

	for i := 0; i < numBits; i++ {
		candidate := flipBit(lastElem, i)
		_, exists := existingValues[candidate]
		if !exists && differsByOne(lastElem, candidate) {
			existingValues[candidate] = true
			result := findGrayCode(numBits, append(currentPath, candidate), existingValues)
			if result != nil {
				return result
			}
			delete(existingValues, candidate)
		}
	}

	return nil
}

func flipBit(value int64, pos int) int64 {
	if pos > 64 {
		panic(fmt.Errorf("unsupported postion %v", pos))
	}
	bit := int64(1) << pos
	return value ^ bit
}

func differsByOne(value1, value2 int64) bool {
	xor := value1 ^ value2
	lastBitZeroed := xor & (xor - 1)
	return xor != 0 && lastBitZeroed == 0
}
