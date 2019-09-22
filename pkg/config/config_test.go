package config

import (
	"testing"
)

var (
	testCfg = GeneratorConfig{
		NumWords:                3,
		WordLenMin:              3,
		WordLenMax:              8,
		CaseTransform:           "ALTERNATE",
		SeparatorCharacter:      "RANDOM",
		SeparatorAlphabet:       DefaultAlphabet,
		PaddingDigitsBefore:     2,
		PaddingDigitsAfter:      2,
		PaddingType:             "FIXED",
		PaddingCharacter:        "RANDOM",
		SymbolAlphabet:          DefaultAlphabet,
		PadToLength:             8,
		PaddingCharactersBefore: 2,
		PaddingCharactersAfter:  2,
	}
)

func TestGeneratorConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		mod     func(*GeneratorConfig)
		wantErr bool
	}{
		{
			name: "valid",
		},
		{
			name: "num words",
			mod: func(gc *GeneratorConfig) {
				gc.NumWords = -1
			},
			wantErr: true,
		},
		{
			name: "word len min",
			mod: func(gc *GeneratorConfig) {
				gc.WordLenMin = -1
			},
			wantErr: true,
		},
		{
			name: "word len max",
			mod: func(gc *GeneratorConfig) {
				gc.WordLenMax = -1
			},
			wantErr: true,
		},
		{
			name: "word len max less than min",
			mod: func(gc *GeneratorConfig) {
				gc.WordLenMax = 1
				gc.WordLenMin = 2
			},
			wantErr: true,
		},
		{
			name: "padding digits before",
			mod: func(gc *GeneratorConfig) {
				gc.PaddingDigitsBefore = -1
			},
			wantErr: true,
		},
		{
			name: "padding digits after",
			mod: func(gc *GeneratorConfig) {
				gc.PaddingDigitsAfter = -1
			},
			wantErr: true,
		},
		{
			name: "pad to length",
			mod: func(gc *GeneratorConfig) {
				gc.PadToLength = -1
			},
			wantErr: true,
		},
		{
			name: "padding characters before",
			mod: func(gc *GeneratorConfig) {
				gc.PaddingCharactersBefore = -1
			},
			wantErr: true,
		},
		{
			name: "padding characters after",
			mod: func(gc *GeneratorConfig) {
				gc.PaddingCharactersAfter = -1
			},
			wantErr: true,
		},
	}
	for _, tc := range tests {
		cfg := testCfg
		if tc.mod != nil {
			tc.mod(&cfg)
		}

		t.Run(tc.name, func(t *testing.T) {
			if err := cfg.Validate(); (err != nil) != tc.wantErr {
				t.Errorf("GeneratorConfig.Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
