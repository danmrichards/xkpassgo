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
	for range before {
		ppw.WriteString(strconv.Itoa(r.Intn(10)))
	}

	ppw.WriteString(pw)

	for range after {
		ppw.WriteString(strconv.Itoa(r.Intn(10)))
	}

	return ppw.String()
}
