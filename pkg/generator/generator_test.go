package generator

import (
	"math/rand"
	"testing"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

func TestXKPassword_Generate(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	tests := []struct {
		name   string
		cfg    *config.GeneratorConfig
		wantPw string
	}{
		{
			name: "appleid",
			cfg: &config.GeneratorConfig{
				NumWords:                3,
				WordLenMin:              5,
				WordLenMax:              7,
				CaseTransform:           "RANDOM",
				SeparatorCharacter:      "RANDOM",
				SeparatorAlphabet:       []string{"-", ":", ".", ","},
				PaddingDigitsBefore:     2,
				PaddingDigitsAfter:      2,
				PaddingType:             "FIXED",
				PaddingCharacter:        "RANDOM",
				SymbolAlphabet:          []string{"!", "?", "@", "&"},
				PaddingCharactersBefore: 1,
				PaddingCharactersAfter:  1,
			},
			wantPw: "!41-JASON-SEWING-CHAIN-29!",
		},
		{
			name: "default",
			cfg: &config.GeneratorConfig{
				NumWords:                3,
				WordLenMin:              4,
				WordLenMax:              8,
				CaseTransform:           "ALTERNATE",
				SeparatorCharacter:      "RANDOM",
				SeparatorAlphabet:       config.DefaultAlphabet,
				PaddingDigitsBefore:     2,
				PaddingDigitsAfter:      2,
				PaddingType:             "FIXED",
				PaddingCharacter:        "RANDOM",
				SymbolAlphabet:          config.DefaultAlphabet,
				PaddingCharactersBefore: 2,
				PaddingCharactersAfter:  2,
			},
			wantPw: "==65:syndrome:PLAYING:such:68==",
		},
		{
			name: "ntlm",
			cfg: &config.GeneratorConfig{
				NumWords:           2,
				WordLenMin:         5,
				WordLenMax:         5,
				CaseTransform:      "INVERT",
				SeparatorCharacter: "RANDOM",
				SeparatorAlphabet: []string{
					"-", "+", "=", ".", "*", "_", "|", "~", ",",
				},
				PaddingDigitsBefore: 1,
				PaddingType:         "FIXED",
				PaddingCharacter:    "RANDOM",
				SymbolAlphabet: []string{
					"!", "@", "$", "%", "^", "&",
					"*", "+", "=", ":", "|", "~", "?",
				},
				PaddingCharactersAfter: 1,
			},
			wantPw: "6.rOUTE.mODES.:",
		},
		{
			name: "securityq",
			cfg: &config.GeneratorConfig{
				NumWords:               6,
				WordLenMin:             4,
				WordLenMax:             8,
				CaseTransform:          "NONE",
				SeparatorCharacter:     " ",
				PaddingType:            "FIXED",
				PaddingCharacter:       "RANDOM",
				SymbolAlphabet:         []string{".", "!", "?"},
				PaddingCharactersAfter: 1,
			},
			wantPw: "equal horror energy cylinder pentium sleeve!",
		},
		{
			name: "web16",
			cfg: &config.GeneratorConfig{
				NumWords:           3,
				WordLenMin:         4,
				WordLenMax:         4,
				CaseTransform:      "RANDOM",
				SeparatorCharacter: "RANDOM",
				SeparatorAlphabet: []string{
					"-", "+", "=", ".", "*", "_", "|", "~", ",",
				},
				PaddingType:      "FIXED",
				PaddingCharacter: "RANDOM",
				SymbolAlphabet: []string{
					"!", "@", "$", "%", "^", "&", "*",
					"+", "=", ":", "|", "~", "?",
				},
				PaddingCharactersBefore: 1,
				PaddingCharactersAfter:  1,
			},
			wantPw: "%held_gulf_tall%",
		},
		{
			name: "web32",
			cfg: &config.GeneratorConfig{
				NumWords:           4,
				WordLenMin:         4,
				WordLenMax:         5,
				CaseTransform:      "ALTERNATE",
				SeparatorCharacter: "RANDOM",
				SeparatorAlphabet: []string{
					"-", "+", "=", ".", "*", "_", "|", "~", ",",
				},
				PaddingDigitsBefore: 2,
				PaddingDigitsAfter:  2,
				PaddingType:         "FIXED",
				PaddingCharacter:    "RANDOM",
				SymbolAlphabet: []string{
					"!", "@", "$", "%", "^", "&", "*", "+", "=", ":", "|", "~", "?",
				},
				PaddingCharactersBefore: 1,
				PaddingCharactersAfter:  1,
			},
			wantPw: "@12_jazz_SAME_views_COLON_86@",
		},
		{
			name: "wifi",
			cfg: &config.GeneratorConfig{
				NumWords:           6,
				WordLenMin:         4,
				WordLenMax:         8,
				CaseTransform:      "RANDOM",
				SeparatorCharacter: "RANDOM",
				SeparatorAlphabet: []string{
					"-", "+", "=", ".", "*", "_", "|", "~", ",",
				},
				PaddingDigitsBefore: 4,
				PaddingDigitsAfter:  4,
				PaddingType:         "ADAPTIVE",
				PaddingCharacter:    "RANDOM",
				SymbolAlphabet: []string{
					"!", "@", "$", "%", "^", "&", "*",
					"+", "=", ":", "|", "~", "?",
				},
				PadToLength: 63,
			},
			wantPw: `0283*reported*preston*yellow*TROOPS*socket*ADAPTOR*1247%%%%%%%%`,
		},
		{
			name: "xkcd",
			cfg: &config.GeneratorConfig{
				NumWords:           4,
				WordLenMin:         4,
				WordLenMax:         8,
				CaseTransform:      "RANDOM",
				SeparatorCharacter: "-",
				PaddingType:        "NONE",
			},
			wantPw: "PLANNED-approved-ANNA-prague",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xkp := &XKPassword{
				r:   r,
				cfg: tc.cfg,
			}

			pw, err := xkp.Generate()
			if err != nil {
				t.Fatalf("Generate error = %v", err)
			}
			if pw != tc.wantPw {
				t.Fatalf("Generate pw = %q, wantPw %q", pw, tc.wantPw)
			}
		})
	}
}
