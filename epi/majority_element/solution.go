package majority_element

func MajoritySearch(stream <-chan string) string {
	candidate := <-stream
	count := 1
	for c := range stream {
		switch {
		case count == 0:
			candidate = c
			count++
		case c != candidate:
			count--
		case c == candidate:
			count++
		}
	}
	return candidate
}
