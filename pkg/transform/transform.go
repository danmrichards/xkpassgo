package transform

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

// ErrInvalidTransform is returned when an invalid case transformation is specified.
var ErrInvalidTransform = errors.New("invalid case transformation")

// transformFunc is a func that transforms the elements of parts and returns
// the transformed elements in a slice.
type transformFunc func(parts []string, r config.Intner) ([]string, error)

// Do returns a slice of parts transformed with style s.
func Do(parts []string, cfg *config.GeneratorConfig, r config.Intner) ([]string, error) {
	tf, ok := styleFuncs[style(cfg.CaseTransform)]
	if !ok {
		return nil, fmt.Errorf("%q: %w", cfg.CaseTransform, ErrInvalidTransform)
	}

	return tf(parts, r)
}

// noop just returns the parts.
func noop(parts []string, _ config.Intner) ([]string, error) {
	return parts, nil
}

// alternate returns "alternating WORD case" parts.
func alternate(parts []string, _ config.Intner) ([]string, error) {
	for i, p := range parts {
		if i%2 == 0 {
			parts[i] = strings.ToLower(p)
		} else {
			parts[i] = strings.ToUpper(p)
		}
	}

	return parts, nil
}

// capitalise returns "Capitalise First Letter" parts.
func capitalise(parts []string, _ config.Intner) ([]string, error) {
	for i, p := range parts {
		if p == "" {
			continue
		}

		r := []rune(p)
		r[0] = unicode.ToTitle(r[0])
		parts[i] = string(r)
	}

	return parts, nil
}

// invert returns "cAPITALISE eVERY lETTER eXCEPT tHE fIRST" parts.
func invert(parts []string, _ config.Intner) ([]string, error) {
	for i := range parts {
		// Strings are immutable so we need to cast to byte slice to be able
		// to modify specific characters.
		pw := []byte(parts[i])
		for j, w := range pw {
			if j == 0 {
				continue
			}

			pw[j] = byte(unicode.ToTitle(rune(w))) //nolint:gosec // G115: ASCII characters fit in byte
		}

		parts[i] = string(pw)
	}

	return parts, nil
}

// lower returns "lower case" parts.
func lower(parts []string, _ config.Intner) ([]string, error) {
	for i, p := range parts {
		parts[i] = strings.ToLower(p)
	}

	return parts, nil
}

// upper returns "UPPER CASE" parts.
func upper(parts []string, _ config.Intner) ([]string, error) {
	for i, p := range parts {
		parts[i] = strings.ToUpper(p)
	}

	return parts, nil
}

// random returns "EVERY word randomly CAPITALISED or NOT" parts.
func random(parts []string, r config.Intner) ([]string, error) {
	for i, p := range parts {
		ri, err := r.Intn(2)
		if err != nil {
			return nil, fmt.Errorf("random case: %w", err)
		}

		if ri == 0 {
			parts[i] = strings.ToUpper(p)
		} else {
			parts[i] = strings.ToLower(p)
		}
	}

	return parts, nil
}
