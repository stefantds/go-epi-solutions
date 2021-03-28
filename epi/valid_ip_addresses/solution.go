package valid_ip_addresses

import (
	"strconv"
)

func GetValidIpAddress(s string) []string {
	return getValidIpAddress(s, 4)
}

func getValidIpAddress(s string, parts int) []string {
	result := make([]string, 0)
	if parts == 0 {
		if s == "" {
			return []string{""}
		}
		return []string{}
	}

	for i := 1; i <= len(s) && i <= 3; i++ {
		prefix, restS := s[:i], s[i:]
		if isValidPart(prefix) {
			remaining := getValidIpAddress(restS, parts-1)
			for _, s := range remaining {
				if s == "" {
					result = append(result, prefix)
				} else {
					result = append(result, prefix+"."+s)
				}
			}
		}
	}

	return result
}

func isValidPart(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '0' && len(s) > 1 {
		return false
	}

	intVal, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if intVal < 0 || intVal >= 256 {
		return false
	}

	return true
}
