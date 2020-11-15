package anagrams

import (
	"sort"
	"strings"
)

func FindAnagrams(dictionary []string) [][]string {
	anagramsDict := make(map[string]*[]string)

	for _, e := range dictionary {
		profile := getProfile(e)
		if _, ok := anagramsDict[profile]; !ok {
			anagramsDict[profile] = &[]string{e}
		} else {
			*(anagramsDict[profile]) = append(*(anagramsDict[profile]), e)
		}
	}

	result := make([][]string, 0)
	for _, a := range anagramsDict {
		if len(*a) >= 2 {
			result = append(result, *a)
		}
	}

	return result
}

func getProfile(s string) string {
	p := sort.StringSlice(strings.Split(s, ""))
	p.Sort()
	return strings.Join(p, "")
}
