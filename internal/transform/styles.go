package transform

// Style represents a style of word transformation.
type Style string

const (
	// None is a no-op transformation and does nothing.
	None Style = "NONE"

	// Alternate is an "alternating WORD case" transformation.
	Alternate Style = "ALTERNATE"

	// Capitalise is a "Capitalise First Letter" transformation.
	Capitalise Style = "CAPITALISE"

	// Invert is a "cAPITALISE eVERY lETTER eXCEPT tHE fIRST" transformation.
	Invert Style = "INVERT"

	// Lower is a "lower case" transformation.
	Lower Style = "LOWER"

	// Upper is an "UPPER CASE" transformation.
	Upper Style = "UPPER"

	// Random is a "EVERY word randomly CAPITALISED or NOT" transformation.
	Random Style = "RANDOM"
)

var styleFuncs = map[Style]transformFunc{
	None:       noop,
	Alternate:  alternate,
	Capitalise: capitalise,
	Invert:     invert,
	Lower:      lower,
	Upper:      upper,
	Random:     random,
}

// Styles returns a list of available transformation styles.
func Styles() []string {
	s := make([]string, 0, len(styleFuncs))
	for style := range styleFuncs {
		s = append(s, string(style))
	}

	return s
}
