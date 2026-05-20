package padding

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

// digits returns pw with the given amount of random digits padded at the
// start and end of the string.
func digits(pw string, cfg *config.GeneratorConfig, r config.Intner) (string, error) {
	before, after := cfg.PaddingDigitsBefore, cfg.PaddingDigitsAfter

	var ppw strings.Builder

	for range before {
		ri, err := r.Intn(10)
		if err != nil {
			return "", fmt.Errorf("padding digits: %w", err)
		}

		ppw.WriteString(strconv.Itoa(ri))
	}

	ppw.WriteString(pw)

	for range after {
		ri, err := r.Intn(10)
		if err != nil {
			return "", fmt.Errorf("padding digits: %w", err)
		}

		ppw.WriteString(strconv.Itoa(ri))
	}

	return ppw.String(), nil
}
