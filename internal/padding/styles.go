package padding

// Style represents a style of symbol padding
type Style string

const (
	// Fixed is a padding style that adds the given number of characters to the
	// start and end of a password. Padding character can be specified or
	// randomly chosen from an alphabet.
	Fixed Style = "FIXED"

	// Adaptive is a padding style that pads a password to a given length. The
	// padding character can be specified or randomly chosen from an alphabet.
	Adaptive Style = "ADAPTIVE"
)

var styleFuncs = map[Style]struct{}{
	Fixed:    struct{}{},
	Adaptive: struct{}{},
}

// Styles returns a list of available padding styles.
func Styles() []string {
	s := make([]string, 0, len(styleFuncs))
	for style := range styleFuncs {
		s = append(s, string(style))
	}

	return s
}
