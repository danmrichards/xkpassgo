package generator

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/danmrichards/xkpassgo/pkg/config"
	"github.com/danmrichards/xkpassgo/pkg/padding"
	"github.com/danmrichards/xkpassgo/pkg/separator"
	"github.com/danmrichards/xkpassgo/pkg/transform"
	"github.com/gobuffalo/packr/v2"
)

// XKPassword is a generator of XKCD-style memorable passwords.
type XKPassword struct {
	r     *rand.Rand
	cfg   *config.GeneratorConfig
	words [][]byte
}

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

	pts, err = transform.Do(pts, xk.cfg, xk.r)
	if err != nil {
		return "", err
	}

	pts, err = separator.Do(pts, xk.cfg, xk.r)
	if err != nil {
		return "", err
	}

	pw = strings.TrimSpace(strings.Join(pts, ""))

	pw, err = padding.Do(pw, xk.cfg, xk.r)
	if err != nil {
		return "", err
	}

	return pw, nil
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
