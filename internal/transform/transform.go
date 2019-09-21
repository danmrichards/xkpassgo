package transform

import (
	"fmt"
	"strings"
)

type transformFunc func(parts []string) []string

func Do(parts []string, s Style) ([]string, error) {
	tf, ok := styleFuncs[s]
	if !ok {
		return nil, fmt.Errorf("%q is not a valid transformation", string(s))
	}

	return tf(parts), nil
}

func noop(parts []string) []string {
	return parts
}

func alternate(parts []string) []string {
	for i, p := range parts {
		if i%2 == 0 {
			parts[i] = strings.ToLower(p)
		} else {
			parts[i] = strings.ToUpper(p)
		}
	}

	return parts
}
