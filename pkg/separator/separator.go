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

	if endPad(cfg) {
		sp = append(sp, char)
	}

	for i, p := range parts {
		sp = append(sp, p)

		if i < len(parts)-1 || endPad(cfg) {
			sp = append(sp, char)
		}
	}

	return sp, nil
}

func endPad(cfg *config.GeneratorConfig) bool {
	padChars := cfg.PaddingCharactersBefore > 0 || cfg.PaddingCharactersAfter > 0

	if cfg.PaddingDigitsBefore == 0 && cfg.PaddingDigitsAfter == 0 && padChars {
		return false
	}

	if cfg.PaddingDigitsBefore > 0 || cfg.PaddingDigitsAfter > 0 {
		return true
	}

	if padChars {
		return true
	}

	return false
}
