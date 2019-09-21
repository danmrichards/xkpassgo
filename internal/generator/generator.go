package generator

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/danmrichards/xkpassgo/internal/config"
	"github.com/danmrichards/xkpassgo/internal/transform"
	"github.com/gobuffalo/packr/v2"
)

type XKPassword struct {
	r     *rand.Rand
	cfg   *config.GeneratorConfig
	words [][]byte
}

func NewXKPassword(cfg *config.GeneratorConfig) *XKPassword {
	return &XKPassword{
		// Create a new pseudo-random source of entropy.
		r:   rand.New(rand.NewSource(time.Now().Unix())),
		cfg: cfg,
	}
}

func (xk *XKPassword) Generate() (string, error) {
	if err := xk.loadWordList(); err != nil {
		return "", err
	}

	pts := xk.parts()

	tpt, err := transform.Do(pts, xk.cfg.CaseTransform)
	if err != nil {
		return "", err
	}

	// TODO: case transformation

	// TODO: seperators

	// TODO: padding digits

	// TODO: padding characters

	return strings.Join(tpt, ""), nil
}

func (xk *XKPassword) loadWordList() error {
	// Load the word file from packed asset.
	box := packr.New("assets", "../../assets")
	wf, err := box.Find("words")
	if err != nil {
		return fmt.Errorf("load words list: %w", err)
	}

	// Split into lines so we can shuffle and select suitable words.
	xk.words = bytes.Split(wf, []byte("\n"))
	return nil
}

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
