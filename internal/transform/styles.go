package transform

type Style string

const (
	None       Style = "NONE"
	Alternate  Style = "ALTERNATE"
	Capitalise Style = "CAPITALISE"
	Invert     Style = "INVERT"
	Lower      Style = "LOWER"
	Upper      Style = "UPPER"
	Random     Style = "RANDOM"
)

var styleFuncs = map[Style]transformFunc{
	None:      noop,
	Alternate: alternate,
	// TODO: Capitalise: capitalise,
	// TODO: Invert:     invert,
	// TODO: Lower:      lower,
	// TODO: Upper:      upper,
	// TODO: Random:     random,
}

func Styles() []string {
	s := make([]string, 0, len(styleFuncs))
	for style := range styleFuncs {
		s = append(s, string(style))
	}

	return s
}
