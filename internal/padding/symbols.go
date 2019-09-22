package padding

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/danmrichards/xkpassgo/internal/config"
)

// Random indicates that a random character from the symbol alphabet should
// be used.
const Random = "RANDOM"

// Symbols returns parts with the configured padding applied.
func Symbols(parts []string, cfg *config.GeneratorConfig, r *rand.Rand) ([]string, error) {
	char := cfg.PaddingCharacter
	alpha := cfg.SymbolAlphabet
	if char == Random {
		char = alpha[r.Intn(len(alpha))]
	}

	switch Style(cfg.PaddingType) {
	case Fixed:
		return fixed(
			parts, char, cfg.PaddingCharactersBefore, cfg.PaddingCharactersAfter,
		), nil
	case Adaptive:
		return adaptive(parts, char, cfg.PadToLength), nil
	default:
		return nil, fmt.Errorf(
			"%q is not a valid transformation", cfg.PaddingType,
		)
	}
}

func fixed(parts []string, char string, before, after int) []string {
	p := make([]string, 0, len(parts)+before+after)

	for i := 0; i < before; i++ {
		p = append(p, char)
	}

	p = append(p, parts...)

	for j := 0; j < after; j++ {
		p = append(p, char)
	}

	return p
}

func adaptive(parts []string, char string, padLen int) []string {
	pwLen := len(strings.Join(parts, ""))

	// Don't attempt to pad if the desired length is less than final length
	// of the password.
	if padLen < pwLen {
		return parts
	}

	diff := padLen - pwLen
	var b strings.Builder
	for i := 0; i < diff; i++ {
		b.WriteString(char)
	}

	parts = append(parts, b.String())
	return parts
}
