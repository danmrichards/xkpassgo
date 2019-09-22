package padding

import (
	"math/rand"
	"strconv"

	"github.com/danmrichards/xkpassgo/internal/config"
)

// Digits returns parts with the given amount of random digits padded at the
// start and end of the slice.
func Digits(parts []string, cfg *config.GeneratorConfig, r *rand.Rand) ([]string, error) {
	before, after := cfg.PaddingDigitsBefore, cfg.PaddingDigitsAfter

	p := make([]string, 0, len(parts)+before+after)

	for i := 0; i < before; i++ {
		p = append(p, strconv.Itoa(r.Intn(10)))
	}

	p = append(p, parts...)

	for j := 0; j < after; j++ {
		p = append(p, strconv.Itoa(r.Intn(10)))
	}

	return p, nil
}
