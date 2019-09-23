package separator

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

var testParts = []string{"correct", "horse", "battery", "staple"}

func TestDo(t *testing.T) {
	r := rand.New(rand.NewSource(1))
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
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parts := make([]string, len(tc.parts))
			copy(parts, tc.parts)

			sp, err := Do(tc.parts, tc.cfg, r)
			if err != nil {
				t.Fatalf("Do error = %v", err)
			}
			if !reflect.DeepEqual(sp, tc.wantParts) {
				t.Errorf("Do parts = %v, wantParts %v", sp, tc.wantParts)
			}
		})
	}
}
