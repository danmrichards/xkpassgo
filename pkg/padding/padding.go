package padding

import (
	"github.com/danmrichards/xkpassgo/pkg/config"
)

// Do returns parts with the configured digit and symbol padding applied.
func Do(pw string, cfg *config.GeneratorConfig, r config.Intner) (string, error) {
	var err error

	pw, err = digits(pw, cfg, r)
	if err != nil {
		return "", err
	}

	return symbols(pw, cfg, r)
}
