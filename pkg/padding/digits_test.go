package padding

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

func TestDigits(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	tests := []struct {
		name   string
		parts  []string
		cfg    *config.GeneratorConfig
		wantPW string
	}{
		{
			name:  "valid same padding",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				PaddingDigitsBefore: 2,
				PaddingDigitsAfter:  2,
			},
			wantPW: "17correcthorsebatterystaple79",
		},
		{
			name:  "valid just before",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				PaddingDigitsBefore: 2,
			},
			wantPW: "18correcthorsebatterystaple",
		},
		{
			name:  "valid just after",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				PaddingDigitsAfter: 2,
			},
			wantPW: "correcthorsebatterystaple50",
		},
		{
			name:   "no padding",
			parts:  testParts,
			cfg:    &config.GeneratorConfig{},
			wantPW: strings.Join(testParts, ""),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parts := make([]string, len(tc.parts))
			copy(parts, tc.parts)

			sp := digits(strings.Join(tc.parts, ""), tc.cfg, r)
			if sp != tc.wantPW {
				t.Errorf("Do parts = %q, wantPW %q", sp, tc.wantPW)
			}
		})
	}
}
