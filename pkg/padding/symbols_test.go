package padding

import (
	"math/rand"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

func TestSymbols(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	tests := []struct {
		name    string
		parts   []string
		cfg     *config.GeneratorConfig
		wantPw  string
		wantErr bool
	}{
		{
			name: "invalid style",
			cfg: &config.GeneratorConfig{
				PaddingType: "foo",
			},
			wantErr: true,
		},
		{
			name: "fixed no padding",
			cfg: &config.GeneratorConfig{
				PaddingType:      string(Fixed),
				PaddingCharacter: "@",
			},
			parts:  testParts,
			wantPw: strings.Join(testParts, ""),
		},
		{
			name: "fixed specific char",
			cfg: &config.GeneratorConfig{
				PaddingType:             string(Fixed),
				PaddingCharacter:        "@",
				PaddingCharactersBefore: 2,
				PaddingCharactersAfter:  2,
			},
			parts:  testParts,
			wantPw: "@@correcthorsebatterystaple@@",
		},
		{
			name: "fixed specific char just before",
			cfg: &config.GeneratorConfig{
				PaddingType:             string(Fixed),
				PaddingCharacter:        "@",
				PaddingCharactersBefore: 2,
			},
			parts:  testParts,
			wantPw: "@@correcthorsebatterystaple",
		},
		{
			name: "fixed specific char just after",
			cfg: &config.GeneratorConfig{
				PaddingType:            string(Fixed),
				PaddingCharacter:       "@",
				PaddingCharactersAfter: 2,
			},
			parts:  testParts,
			wantPw: "correcthorsebatterystaple@@",
		},
		{
			name: "fixed random char default alpha",
			cfg: &config.GeneratorConfig{
				PaddingType:             string(Fixed),
				PaddingCharacter:        string(Random),
				SymbolAlphabet:          config.DefaultAlphabet,
				PaddingCharactersBefore: 2,
				PaddingCharactersAfter:  2,
			},
			parts:  testParts,
			wantPw: "&&correcthorsebatterystaple&&",
		},
		{
			name: "fixed random char default alpha just before",
			cfg: &config.GeneratorConfig{
				PaddingType:             string(Fixed),
				PaddingCharacter:        string(Random),
				SymbolAlphabet:          config.DefaultAlphabet,
				PaddingCharactersBefore: 2,
			},
			parts:  testParts,
			wantPw: "//correcthorsebatterystaple",
		},
		{
			name: "fixed random char default alpha just after",
			cfg: &config.GeneratorConfig{
				PaddingType:            string(Fixed),
				PaddingCharacter:       string(Random),
				SymbolAlphabet:         config.DefaultAlphabet,
				PaddingCharactersAfter: 2,
			},
			parts:  testParts,
			wantPw: "correcthorsebatterystaple::",
		},
		{
			name: "fixed random char specific alpha",
			cfg: &config.GeneratorConfig{
				PaddingType:             string(Fixed),
				PaddingCharacter:        string(Random),
				SymbolAlphabet:          testAlpha,
				PaddingCharactersBefore: 2,
				PaddingCharactersAfter:  2,
			},
			parts:  testParts,
			wantPw: "$$correcthorsebatterystaple$$",
		},
		{
			name: "fixed random char specific alpha just before",
			cfg: &config.GeneratorConfig{
				PaddingType:             string(Fixed),
				PaddingCharacter:        string(Random),
				SymbolAlphabet:          testAlpha,
				PaddingCharactersBefore: 2,
			},
			parts:  testParts,
			wantPw: "@@correcthorsebatterystaple",
		},
		{
			name: "fixed random char specific alpha just after",
			cfg: &config.GeneratorConfig{
				PaddingType:            string(Fixed),
				PaddingCharacter:       string(Random),
				SymbolAlphabet:         testAlpha,
				PaddingCharactersAfter: 2,
			},
			parts:  testParts,
			wantPw: "correcthorsebatterystaple££",
		},
		{
			name: "adaptive no padding",
			cfg: &config.GeneratorConfig{
				PaddingType:      string(Adaptive),
				PaddingCharacter: "@",
			},
			parts:  testParts,
			wantPw: strings.Join(testParts, ""),
		},
		{
			name: "adaptive pad less than length",
			cfg: &config.GeneratorConfig{
				PaddingType:      string(Adaptive),
				PaddingCharacter: "@",
				PadToLength:      20,
			},
			parts:  testParts,
			wantPw: strings.Join(testParts, ""),
		},
		{
			name: "adaptive specific char",
			cfg: &config.GeneratorConfig{
				PaddingType:      string(Adaptive),
				PaddingCharacter: "£",
				PadToLength:      40,
			},
			parts:  testParts,
			wantPw: "correcthorsebatterystaple£££££££££££££££",
		},
		{
			name: "adaptive random char default alpha",
			cfg: &config.GeneratorConfig{
				PaddingType:      string(Adaptive),
				PaddingCharacter: string(Random),
				SymbolAlphabet:   config.DefaultAlphabet,
				PadToLength:      40,
			},
			parts:  testParts,
			wantPw: "correcthorsebatterystaple---------------",
		},
		{
			name: "adaptive random char specific alpha",
			cfg: &config.GeneratorConfig{
				PaddingType:      string(Adaptive),
				PaddingCharacter: string(Random),
				SymbolAlphabet:   testAlpha,
				PadToLength:      40,
			},
			parts:  testParts,
			wantPw: "correcthorsebatterystaple!!!!!!!!!!!!!!!",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parts := make([]string, len(tc.parts))
			copy(parts, tc.parts)

			p, err := symbols(strings.Join(parts, ""), tc.cfg, r)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Do error = %v, wantErr %v", err, tc.wantErr)
			}
			if p != tc.wantPw {
				t.Errorf("Do parts = %q, wantPw %q", p, tc.wantPw)
			}

			if tc.cfg.PadToLength > 0 && tc.cfg.PadToLength > len(strings.Join(tc.parts, "")) {
				if pLen := utf8.RuneCountInString(p); pLen != tc.cfg.PadToLength {
					t.Errorf(
						"Do pad to length = %d, got %d (%q)",
						tc.cfg.PadToLength, pLen, p,
					)
				}
			}
		})
	}
}
