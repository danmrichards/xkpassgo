package padding

import (
	"math/rand"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

// Do returns parts with the configured digit and symbol padding applied.
func Do(pw string, cfg *config.GeneratorConfig, r *rand.Rand) (string, error) {
	pw = digits(pw, cfg, r)

	return symbols(pw, cfg, r)
}
