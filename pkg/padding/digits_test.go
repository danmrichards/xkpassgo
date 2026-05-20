package padding

import (
	mathrand "math/rand"
	"strings"
	"testing"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

var testDigitsRand = &syncIntner{r: mathrand.New(mathrand.NewSource(1))}

func TestDigits(t *testing.T) { //nolint:paralleltest
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

	for _, tc := range tests { //nolint:paralleltest
		t.Run(tc.name, func(t *testing.T) {
			parts := make([]string, len(tc.parts))
			copy(parts, tc.parts)

			sp, err := digits(strings.Join(tc.parts, ""), tc.cfg, testDigitsRand)
			if err != nil {
				t.Fatalf("digits error = %v", err)
			}

			if sp != tc.wantPW {
				t.Errorf("Do parts = %q, wantPW %q", sp, tc.wantPW)
			}
		})
	}
}
