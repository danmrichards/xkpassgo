package transform

// style represents a style of word transformation.
type style string

const (
	// None is a no-op transformation and does nothing.
	None style = "NONE"

	// Alternate is an "alternating WORD case" transformation.
	Alternate style = "ALTERNATE"

	// Capitalise is a "Capitalise First Letter" transformation.
	Capitalise style = "CAPITALISE"

	// Invert is a "cAPITALISE eVERY lETTER eXCEPT tHE fIRST" transformation.
	Invert style = "INVERT"

	// Lower is a "lower case" transformation.
	Lower style = "LOWER"

	// Upper is an "UPPER CASE" transformation.
	Upper style = "UPPER"

	// Random is a "EVERY word randomly CAPITALISED or NOT" transformation.
	Random style = "RANDOM"
)

var styleFuncs = map[style]transformFunc{
	None:       noop,
	Alternate:  alternate,
	Capitalise: capitalise,
	Invert:     invert,
	Lower:      lower,
	Upper:      upper,
	Random:     random,
}
