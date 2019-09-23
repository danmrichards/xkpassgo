package padding

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

// digits returns pw with the given amount of random digits padded at the
// start and end of the string.
func digits(pw string, cfg *config.GeneratorConfig, r *rand.Rand) string {
	before, after := cfg.PaddingDigitsBefore, cfg.PaddingDigitsAfter

	var ppw strings.Builder
	for i := 0; i < before; i++ {
		ppw.WriteString(strconv.Itoa(r.Intn(10)))
	}

	ppw.WriteString(pw)

	for j := 0; j < after; j++ {
		ppw.WriteString(strconv.Itoa(r.Intn(10)))
	}

	return ppw.String()
}
