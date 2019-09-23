package padding

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode/utf8"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

// Random indicates that a random character from the symbol alphabet should
// be used.
const Random = "RANDOM"

// symbols returns pw with the configured padding applied.
func symbols(pw string, cfg *config.GeneratorConfig, r *rand.Rand) (string, error) {
	char := cfg.PaddingCharacter
	alpha := cfg.SymbolAlphabet
	if char == Random {
		char = alpha[r.Intn(len(alpha))]
	}

	switch Style(cfg.PaddingType) {
	case None:
		return pw, nil
	case Fixed:
		return fixed(
			pw, char, cfg.PaddingCharactersBefore, cfg.PaddingCharactersAfter,
		), nil
	case Adaptive:
		return adaptive(pw, char, cfg.PadToLength), nil
	default:
		return "", fmt.Errorf(
			"%q is not a valid padding style", cfg.PaddingType,
		)
	}
}

func fixed(pw string, char string, before, after int) string {
	var ppw strings.Builder

	for i := 0; i < before; i++ {
		ppw.WriteString(char)
	}

	ppw.WriteString(pw)

	for j := 0; j < after; j++ {
		ppw.WriteString(char)
	}

	return ppw.String()
}

func adaptive(pw string, char string, padLen int) string {
	pwLen := utf8.RuneCountInString(pw)

	// Don't attempt to pad if the desired length is less than final length
	// of the password.
	if padLen < pwLen {
		return pw
	}

	diff := padLen - pwLen

	var ppw strings.Builder
	ppw.WriteString(pw)

	for i := 0; i < diff; i++ {
		ppw.WriteString(char)
	}

	return ppw.String()
}
