package separator

import (
	"errors"
	"math/rand"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

// Random indicates that a random character from the separator alphabet should
// be used.
const Random = "RANDOM"

// Do returns parts with the given separator character inserted in between the
// elements.
//
// The separatator character is either literally char or if set to "RANDOM" a
// random character from the given alphabet.
//
// Separators are not added to the start or end of the slice.
func Do(parts []string, cfg *config.GeneratorConfig, r *rand.Rand) ([]string, error) {
	char := cfg.SeparatorCharacter
	alpha := cfg.SeparatorAlphabet

	if char == Random {
		if len(alpha) == 0 {
			return nil, errors.New("configured alphabet cannot be empty")
		}

		char = alpha[r.Intn(len(alpha))]
	}

	sp := make([]string, 0, len(parts)+len(parts)-1)
	for i, p := range parts {
		sp = append(sp, p)

		if i < len(parts)-1 {
			sp = append(sp, char)
		}
	}

	return sp, nil
}
