package string_transformability

import "github.com/stefantds/go-epi-judge/utils"

func TransformString(d map[string]struct{}, s string, t string) int {
	// visited keeps visited words as keys and the distance from s to it as value
	visited := make(map[string]int)
	visited[s] = 0

	bfsQueue := make(utils.Queue, 0)
	bfsQueue = bfsQueue.Enqueue(s)

	for !bfsQueue.IsEmpty() {
		var value interface{}
		bfsQueue, value = bfsQueue.Dequeue()
		current := value.(string)
		currentDistance := visited[current]

		if current == t {
			return currentDistance
		}

		for i := range current {
			for j := 'a'; j <= 'z'; j++ {
				word := []rune(current)
				word[i] = j
				wordStr := string(word)
				if _, exists := d[wordStr]; !exists { // the word is not in the dictionary
					continue
				}
				if _, ok := visited[wordStr]; !ok { // if not already visited
					visited[wordStr] = currentDistance + 1
					bfsQueue = bfsQueue.Enqueue(wordStr)
				}
			}
		}
	}

	return -1
}
