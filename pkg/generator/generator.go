package generator

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/danmrichards/xkpassgo/pkg/assets"
	"github.com/danmrichards/xkpassgo/pkg/config"
	"github.com/danmrichards/xkpassgo/pkg/padding"
	"github.com/danmrichards/xkpassgo/pkg/separator"
	"github.com/danmrichards/xkpassgo/pkg/transform"
)

// XKPassword is a generator of XKCD-style memorable passwords.
type XKPassword struct {
	r     config.Intner
	cfg   *config.GeneratorConfig
	words [][]byte
}

// NewXKPassword returns a new configured XKCD password generator.
func NewXKPassword(cfg *config.GeneratorConfig) *XKPassword {
	return &XKPassword{
		r:   &cryptoRand{},
		cfg: cfg,
	}
}

// Generate returns a new generated password.
func (xk *XKPassword) Generate() (string, error) {
	xk.loadWordList()

	pts, err := xk.parts()
	if err != nil {
		return "", fmt.Errorf("parts: %w", err)
	}

	pts, err = transform.Do(pts, xk.cfg, xk.r)
	if err != nil {
		return "", fmt.Errorf("transform: %w", err)
	}

	pts, err = separator.Do(pts, xk.cfg, xk.r)
	if err != nil {
		return "", fmt.Errorf("separator: %w", err)
	}

	pw := strings.TrimSpace(strings.Join(pts, ""))

	pw, err = padding.Do(pw, xk.cfg, xk.r)
	if err != nil {
		return "", fmt.Errorf("padding: %w", err)
	}

	return pw, nil
}

// loadWordList loads the list of words for generating passwords.
//
// The word list is loaded from an embedded asset.
func (xk *XKPassword) loadWordList() {
	// Split into lines so we can shuffle and select suitable words.
	xk.words = bytes.Split(assets.Words, []byte("\n"))
}

// parts returns a slice of words to use in the generated password.
//
// The number of words in the slice, and the length of those words, is based on
// the configuration of the password generator.
func (xk *XKPassword) parts() ([]string, error) {
	p := make([]string, 0, xk.cfg.NumWords)
	for {
		if len(p) == xk.cfg.NumWords {
			break
		}

		ri, err := xk.r.Intn(len(xk.words))
		if err != nil {
			return nil, fmt.Errorf("random word: %w", err)
		}

		rw := string(xk.words[ri])
		if rwl := len(rw); rwl < xk.cfg.WordLenMin || rwl > xk.cfg.WordLenMax {
			continue
		}

		p = append(p, rw)
	}

	return p, nil
}
