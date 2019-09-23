package padding

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

func TestSymbols(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	tests := []struct {
		name      string
		parts     []string
		cfg       *config.GeneratorConfig
		wantParts []string
		wantErr   bool
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
				PaddingType:      "FIXED",
				PaddingCharacter: "@",
			},
			parts:     testParts,
			wantParts: testParts,
		},
		{
			name: "fixed specific char",
			cfg: &config.GeneratorConfig{
				PaddingType:             "FIXED",
				PaddingCharacter:        "@",
				PaddingCharactersBefore: 2,
				PaddingCharactersAfter:  2,
			},
			parts: testParts,
			wantParts: []string{
				"@", "@", "correct", "horse", "battery", "staple", "@", "@",
			},
		},
		{
			name: "fixed specific char just before",
			cfg: &config.GeneratorConfig{
				PaddingType:             "FIXED",
				PaddingCharacter:        "@",
				PaddingCharactersBefore: 2,
			},
			parts: testParts,
			wantParts: []string{
				"@", "@", "correct", "horse", "battery", "staple",
			},
		},
		{
			name: "fixed specific char just after",
			cfg: &config.GeneratorConfig{
				PaddingType:            "FIXED",
				PaddingCharacter:       "@",
				PaddingCharactersAfter: 2,
			},
			parts: testParts,
			wantParts: []string{
				"correct", "horse", "battery", "staple", "@", "@",
			},
		},
		{
			name: "fixed random char default alpha",
			cfg: &config.GeneratorConfig{
				PaddingType:             "FIXED",
				PaddingCharacter:        "RANDOM",
				SymbolAlphabet:          config.DefaultAlphabet,
				PaddingCharactersBefore: 2,
				PaddingCharactersAfter:  2,
			},
			parts: testParts,
			wantParts: []string{
				"&", "&", "correct", "horse", "battery", "staple", "&", "&",
			},
		},
		{
			name: "fixed random char default alpha just before",
			cfg: &config.GeneratorConfig{
				PaddingType:             "FIXED",
				PaddingCharacter:        "RANDOM",
				SymbolAlphabet:          config.DefaultAlphabet,
				PaddingCharactersBefore: 2,
			},
			parts: testParts,
			wantParts: []string{
				"/", "/", "correct", "horse", "battery", "staple",
			},
		},
		{
			name: "fixed random char default alpha just after",
			cfg: &config.GeneratorConfig{
				PaddingType:            "FIXED",
				PaddingCharacter:       "RANDOM",
				SymbolAlphabet:         config.DefaultAlphabet,
				PaddingCharactersAfter: 2,
			},
			parts: testParts,
			wantParts: []string{
				"correct", "horse", "battery", "staple", ":", ":",
			},
		},
		{
			name: "fixed random char specific alpha",
			cfg: &config.GeneratorConfig{
				PaddingType:             "FIXED",
				PaddingCharacter:        "RANDOM",
				SymbolAlphabet:          testAlpha,
				PaddingCharactersBefore: 2,
				PaddingCharactersAfter:  2,
			},
			parts: testParts,
			wantParts: []string{
				"$", "$", "correct", "horse", "battery", "staple", "$", "$",
			},
		},
		{
			name: "fixed random char specific alpha just before",
			cfg: &config.GeneratorConfig{
				PaddingType:             "FIXED",
				PaddingCharacter:        "RANDOM",
				SymbolAlphabet:          testAlpha,
				PaddingCharactersBefore: 2,
			},
			parts: testParts,
			wantParts: []string{
				"@", "@", "correct", "horse", "battery", "staple",
			},
		},
		{
			name: "fixed random char specific alpha just after",
			cfg: &config.GeneratorConfig{
				PaddingType:            "FIXED",
				PaddingCharacter:       "RANDOM",
				SymbolAlphabet:         testAlpha,
				PaddingCharactersAfter: 2,
			},
			parts: testParts,
			wantParts: []string{
				"correct", "horse", "battery", "staple", "£", "£",
			},
		},
		{
			name: "adaptive no padding",
			cfg: &config.GeneratorConfig{
				PaddingType:      "ADAPTIVE",
				PaddingCharacter: "@",
			},
			parts:     testParts,
			wantParts: testParts,
		},
		{
			name: "adaptive pad less than length",
			cfg: &config.GeneratorConfig{
				PaddingType:      "ADAPTIVE",
				PaddingCharacter: "@",
				PadToLength:      20,
			},
			parts:     testParts,
			wantParts: testParts,
		},
		{
			name: "adaptive specific char",
			cfg: &config.GeneratorConfig{
				PaddingType:      "ADAPTIVE",
				PaddingCharacter: "£",
				PadToLength:      40,
			},
			parts: testParts,
			wantParts: []string{
				"correct", "horse", "battery", "staple", "£££££££££££££££",
			},
		},
		{
			name: "adaptive random char default alpha",
			cfg: &config.GeneratorConfig{
				PaddingType:      "ADAPTIVE",
				PaddingCharacter: "RANDOM",
				SymbolAlphabet:   config.DefaultAlphabet,
				PadToLength:      40,
			},
			parts: testParts,
			wantParts: []string{
				"correct", "horse", "battery", "staple", "---------------",
			},
		},
		{
			name: "adaptive random char specific alpha",
			cfg: &config.GeneratorConfig{
				PaddingType:      "ADAPTIVE",
				PaddingCharacter: "RANDOM",
				SymbolAlphabet:   testAlpha,
				PadToLength:      40,
			},
			parts: testParts,
			wantParts: []string{
				"correct", "horse", "battery", "staple", "!!!!!!!!!!!!!!!",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parts := make([]string, len(tc.parts))
			copy(parts, tc.parts)

			p, err := Symbols(parts, tc.cfg, r)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Do error = %v, wantErr %v", err, tc.wantErr)
			}
			if !reflect.DeepEqual(p, tc.wantParts) {
				t.Errorf("Do parts = %v, wantParts %v", p, tc.wantParts)
			}
		})
	}
}
