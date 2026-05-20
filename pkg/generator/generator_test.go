package generator

import (
	mathrand "math/rand"
	"testing"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

type testIntner struct {
	r *mathrand.Rand
}

func (t testIntner) Intn(n int) (int, error) {
	return t.r.Intn(n), nil
}

var testRand = testIntner{r: mathrand.New(mathrand.NewSource(1))}

func TestXKPassword_Generate(t *testing.T) {
	t.Parallel()

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
			wantPw: "!41-jason-SEWING-CHAIN-29!",
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
			wantPw: "%HELD_gulf_TALL%",
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
			wantPw: `0283*reported*preston*yellow*troops*SOCKET*ADAPTOR*1247%%%%%%%%`,
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
			wantPw: "planned-approved-ANNA-PRAGUE",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Cannot call t.Parallel() here because subtests share the same
			// Rand instance which is not thread-safe.
			xkp := &XKPassword{
				r:   testRand,
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
