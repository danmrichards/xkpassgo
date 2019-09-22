package generator

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/danmrichards/xkpassgo/internal/config"
	"github.com/danmrichards/xkpassgo/internal/padding"
	"github.com/danmrichards/xkpassgo/internal/separator"
	"github.com/danmrichards/xkpassgo/internal/transform"
	"github.com/gobuffalo/packr/v2"
)

// modFunc is a function that applies a modification the password parts.
//
// The func can rely on the given config and random source.
type modFunc func(parts []string, cfg *config.GeneratorConfig, r *rand.Rand) ([]string, error)

// XKPassword is a generator of XKCD-style memorable passwords.
type XKPassword struct {
	r     *rand.Rand
	cfg   *config.GeneratorConfig
	words [][]byte
}

var (
	// mods is the set of password modifications.
	mods = []modFunc{
		transform.Do,
		separator.Do,
		padding.Digits,
		padding.Symbols,
	}
)

// NewXKPassword returns a new configured XKCD password generator.
func NewXKPassword(cfg *config.GeneratorConfig) *XKPassword {
	return &XKPassword{
		// Create a new pseudo-random source of entropy.
		r:   rand.New(rand.NewSource(time.Now().Unix())),
		cfg: cfg,
	}
}

// Generate returns a new generated password.
func (xk *XKPassword) Generate() (pw string, err error) {
	if err = xk.loadWordList(); err != nil {
		return "", err
	}

	pts := xk.parts()

	// Apply each modification to the password parts.
	for _, m := range mods {
		pts, err = m(pts, xk.cfg, xk.r)
		if err != nil {
			return "", err
		}
	}

	return strings.Join(pts, ""), nil
}

// loadWordList loads the list of words for generating passwords.
//
// The word list is loaded from a packed asset file.
func (xk *XKPassword) loadWordList() error {
	box := packr.New("assets", "../../assets")
	wf, err := box.Find("words")
	if err != nil {
		return fmt.Errorf("load words list: %w", err)
	}

	// Split into lines so we can shuffle and select suitable words.
	xk.words = bytes.Split(wf, []byte("\n"))
	return nil
}

// parts returns a slice of words to use in the generated password.
//
// The number of words in the slice, and the length of those words, is based on
// the configuration of the password generator.
func (xk *XKPassword) parts() (p []string) {
	p = make([]string, 0, xk.cfg.NumWords)
	for {
		if len(p) == xk.cfg.NumWords {
			break
		}

		// Get a random word from the list and ensure it meets requirements.
		rw := string(xk.words[xk.r.Intn(len(xk.words))])
		if rwl := len(rw); rwl < xk.cfg.WordLenMin || rwl > xk.cfg.WordLenMax {
			continue
		}

		p = append(p, rw)
	}

	return p
}
