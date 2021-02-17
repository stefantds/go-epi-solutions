package directory_path_normalization

import (
	"fmt"
	"strings"

	"github.com/stefantds/go-epi-judge/data_structures/stack"
)

func ShortestEquivalentPath(path string) string {
	pathParts := strings.Split(path, "/")
	isAbsolute := path[0] == '/'

	s := make(stack.Stack, 0)

	for _, part := range pathParts {
		switch {
		case part == "",
			part == ".": // do nothing, "." is redundant
			break
		case part == ".." && s.IsEmpty() && isAbsolute:
			panic(fmt.Errorf("trying to go up the root: %s", path))
		case part == ".." && s.IsEmpty(), // relative path with ".." in the first position
			part == ".." && !s.IsEmpty() && s.Peek().(string) == "..": // relative path with other ".." parts
			s = s.Push(part)
		case part == ".." && !s.IsEmpty(): // last element is a directory name
			s, _ = s.Pop()
		default:
			s = s.Push(part)
		}
	}

	return pathArrayToString(s, isAbsolute)
}

func pathArrayToString(parts stack.Stack, isAbsolute bool) string {
	var resultBuilder strings.Builder

	if isAbsolute {
		resultBuilder.WriteRune('/')
	}

	for i, p := range parts {
		resultBuilder.WriteString(p.(string))
		if i < len(parts)-1 {
			resultBuilder.WriteRune('/')
		}
	}

	return resultBuilder.String()
}
