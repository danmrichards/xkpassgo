package transform

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/danmrichards/xkpassgo/pkg/config"
)

var (
	testParts        = []string{"correct", "horse", "battery", "staple"}
	testExtremeParts = []string{"cOrReCt", "hOrSe", "bAtTeRy", "sTaPle"}
)

func TestDo(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	tests := []struct {
		name      string
		parts     []string
		cfg       *config.GeneratorConfig
		wantParts []string
		wantErr   bool
	}{
		{
			name: "invaid style",
			cfg: &config.GeneratorConfig{
				CaseTransform: "foo",
			},
			wantErr: true,
		},
		{
			name:  "noop",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				CaseTransform: string(None),
			},
			wantParts: testParts,
		},
		{
			name:  "alternate",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				CaseTransform: string(Alternate),
			},
			wantParts: []string{"correct", "HORSE", "battery", "STAPLE"},
		},
		{
			name:  "alternate extreme",
			parts: testExtremeParts,
			cfg: &config.GeneratorConfig{
				CaseTransform: string(Alternate),
			},
			wantParts: []string{"correct", "HORSE", "battery", "STAPLE"},
		},
		{
			name:  "capitalise",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				CaseTransform: string(Capitalise),
			},
			wantParts: []string{"Correct", "Horse", "Battery", "Staple"},
		},
		{
			name:  "capitalise extreme",
			parts: testExtremeParts,
			cfg: &config.GeneratorConfig{
				CaseTransform: string(Capitalise),
			},
			wantParts: []string{"COrReCt", "HOrSe", "BAtTeRy", "STaPle"},
		},
		{
			name:  "invert",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				CaseTransform: string(Invert),
			},
			wantParts: []string{"cORRECT", "hORSE", "bATTERY", "sTAPLE"},
		},
		{
			name:  "invert extreme",
			parts: testExtremeParts,
			cfg: &config.GeneratorConfig{
				CaseTransform: string(Invert),
			},
			wantParts: []string{"cORRECT", "hORSE", "bATTERY", "sTAPLE"},
		},
		{
			name:  "lower",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				CaseTransform: string(Lower),
			},
			wantParts: testParts,
		},
		{
			name:  "lower extreme",
			parts: testExtremeParts,
			cfg: &config.GeneratorConfig{
				CaseTransform: string(Lower),
			},
			wantParts: testParts,
		},
		{
			name:  "upper",
			parts: testParts,
			cfg: &config.GeneratorConfig{
				CaseTransform: string(Upper),
			},
			wantParts: []string{"CORRECT", "HORSE", "BATTERY", "STAPLE"},
		},
		{
			name:  "upper extreme",
			parts: testExtremeParts,
			cfg: &config.GeneratorConfig{
				CaseTransform: string(Upper),
			},
			wantParts: []string{"CORRECT", "HORSE", "BATTERY", "STAPLE"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parts := make([]string, len(tc.parts))
			copy(parts, tc.parts)

			p, err := Do(parts, tc.cfg, r)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Do error = %v, wantErr %v", err, tc.wantErr)
			}
			if !reflect.DeepEqual(p, tc.wantParts) {
				t.Errorf("Do parts = %v, wantParts %v", p, tc.wantParts)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	tests := []struct {
		name  string
		parts []string
	}{
		{
			name:  "normal",
			parts: testParts,
		},
		{
			name:  "extreme",
			parts: testExtremeParts,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parts := make([]string, len(tc.parts))
			copy(parts, tc.parts)

			rp := random(parts, r)
			if reflect.DeepEqual(rp, testParts) {
				t.Fatalf("Do random = %v: not transformed", rp)
			}

			var hasLower, hasUpper bool
			for _, p := range rp {
				if strings.ToLower(p) == p {
					hasLower = true
				}
				if strings.ToUpper(p) == p {
					hasUpper = true
				}
			}

			if !hasLower {
				t.Errorf("Do random = %v: no lower case words", rp)
			}
			if !hasUpper {
				t.Errorf("Do random = %v: no upper case words", rp)
			}
		})
	}
}
