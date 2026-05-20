package padding

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

// Random indicates that a random character from the symbol alphabet should
// be used.
const Random = "RANDOM"

// ErrInvalidPaddingStyle is returned when an invalid padding style is specified.
var ErrInvalidPaddingStyle = errors.New("invalid padding style")

// symbols returns pw with the configured padding applied.
func symbols(pw string, cfg *config.GeneratorConfig, r config.Intner) (string, error) {
	char := cfg.PaddingCharacter
	alpha := cfg.SymbolAlphabet

	if char == Random {
		ri, err := r.Intn(len(alpha))
		if err != nil {
			return "", fmt.Errorf("padding symbol: %w", err)
		}

		char = alpha[ri]
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
		return "", fmt.Errorf("%q: %w", cfg.PaddingType, ErrInvalidPaddingStyle)
	}
}

func fixed(pw string, char string, before, after int) string {
	var ppw strings.Builder

	for range before {
		ppw.WriteString(char)
	}

	ppw.WriteString(pw)

	for range after {
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

	for range diff {
		ppw.WriteString(char)
	}

	return ppw.String()
}
