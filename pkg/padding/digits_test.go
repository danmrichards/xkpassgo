package padding

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

func TestDigits(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	tests := []struct {
		name      string
		parts     []string
		cfg       *config.GeneratorConfig
		wantParts []string
	}{
		{
			name:  "valid same padding",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				PaddingDigitsBefore: 2,
				PaddingDigitsAfter:  2,
			},
			wantParts: []string{
				"1", "7", "correct", "horse", "battery", "staple", "7", "9",
			},
		},
		{
			name:  "valid just before",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				PaddingDigitsBefore: 2,
			},
			wantParts: []string{
				"1", "8", "correct", "horse", "battery", "staple",
			},
		},
		{
			name:  "valid just after",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				PaddingDigitsAfter: 2,
			},
			wantParts: []string{
				"correct", "horse", "battery", "staple", "5", "0",
			},
		},
		{
			name:      "no padding",
			parts:     testParts,
			cfg:       &config.GeneratorConfig{},
			wantParts: testParts,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parts := make([]string, len(tc.parts))
			copy(parts, tc.parts)

			sp, err := Digits(tc.parts, tc.cfg, r)
			if err != nil {
				t.Fatalf("Do error = %v", err)
			}
			if !reflect.DeepEqual(sp, tc.wantParts) {
				t.Errorf("Do parts = %v, wantParts %v", sp, tc.wantParts)
			}
		})
	}
}
