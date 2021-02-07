package nearest_repeated_entries

import "math"

func FindNearestRepetition(paragraph []string) int {
	lastSeen := make(map[string]int)
	var minDist = math.MaxInt32

	for i, w := range paragraph {
		if prevPos, ok := lastSeen[w]; ok {
			dist := i - prevPos
			if dist < minDist {
				minDist = dist
			}
		}
		lastSeen[w] = i
	}

	if minDist == math.MaxInt32 {
		return -1
	}
	return minDist
}
