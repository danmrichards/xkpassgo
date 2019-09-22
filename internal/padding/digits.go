package padding

import (
	"math/rand"
	"strconv"
)

// Digits returns parts with the given amount of random digits padded at the
// start and end of the slice.
func Digits(parts []string, before, after int, r *rand.Rand) []string {
	p := make([]string, 0, len(parts)+before+after)

	for i := 0; i < before; i++ {
		p = append(p, strconv.Itoa(r.Intn(10)))
	}

	p = append(p, parts...)

	for j := 0; j < after; j++ {
		p = append(p, strconv.Itoa(r.Intn(10)))
	}

	return p
}
