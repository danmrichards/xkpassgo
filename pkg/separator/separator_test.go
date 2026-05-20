package separator

import (
	mathrand "math/rand"
	"reflect"
	"sync"
	"testing"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

type syncIntner struct {
	mu sync.Mutex
	r  *mathrand.Rand
}

func (s *syncIntner) Intn(n int) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.r.Intn(n), nil
}

var testParts = []string{"correct", "horse", "battery", "staple"}

var testSeparatorRand = &syncIntner{r: mathrand.New(mathrand.NewSource(1))}

func TestDo(t *testing.T) { //nolint:paralleltest
	tests := []struct {
		name      string
		parts     []string
		cfg       *config.GeneratorConfig
		wantParts []string
	}{
		{
			name:  "specific char",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				SeparatorCharacter: "@",
			},
			wantParts: []string{
				"correct", "@", "horse", "@", "battery", "@", "staple",
			},
		},
		{
			name:  "specific char pad end",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				SeparatorCharacter:  "@",
				PaddingDigitsBefore: 1,
			},
			wantParts: []string{
				"@", "correct", "@", "horse", "@", "battery", "@", "staple", "@",
			},
		},
		{
			name:  "random char single alpha",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				SeparatorCharacter: Random,
				SeparatorAlphabet:  []string{"@"},
			},
			wantParts: []string{
				"correct", "@", "horse", "@", "battery", "@", "staple",
			},
		},
		{
			name:  "random char single alpha pad end",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				SeparatorCharacter:  Random,
				SeparatorAlphabet:   []string{"@"},
				PaddingDigitsBefore: 1,
			},
			wantParts: []string{
				"@", "correct", "@", "horse", "@", "battery", "@", "staple", "@",
			},
		},
		{
			name:  "random char default alpha",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				SeparatorCharacter: Random,
				SeparatorAlphabet:  config.DefaultAlphabet,
			},
			wantParts: []string{
				"correct", ":", "horse", ":", "battery", ":", "staple",
			},
		},
		{
			name:  "random char default alpha pad end",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				SeparatorCharacter:  Random,
				SeparatorAlphabet:   config.DefaultAlphabet,
				PaddingDigitsBefore: 1,
			},
			wantParts: []string{
				":", "correct", ":", "horse", ":", "battery", ":", "staple", ":",
			},
		},
		{
			name:  "random char specific alpha",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				SeparatorCharacter: Random,
				SeparatorAlphabet:  []string{"!", "@", "£", "$"},
			},
			wantParts: []string{
				"correct", "@", "horse", "@", "battery", "@", "staple",
			},
		},
		{
			name:  "random char specific alpha pad end",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				SeparatorCharacter:  Random,
				SeparatorAlphabet:   []string{"!", "@", "£", "$"},
				PaddingDigitsBefore: 1,
			},
			wantParts: []string{
				"£", "correct", "£", "horse", "£", "battery", "£", "staple", "£",
			},
		},
	}

	for _, tc := range tests { //nolint:paralleltest
		t.Run(tc.name, func(t *testing.T) {
			parts := make([]string, len(tc.parts))
			copy(parts, tc.parts)

			sp, err := Do(tc.parts, tc.cfg, testSeparatorRand)
			if err != nil {
				t.Fatalf("Do error = %v", err)
			}

			if !reflect.DeepEqual(sp, tc.wantParts) {
				t.Errorf("Do parts = %v, wantParts %v", sp, tc.wantParts)
			}
		})
	}
}
