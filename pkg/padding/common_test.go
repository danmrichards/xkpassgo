package padding

import mathrand "math/rand"

type testIntner struct {
	r *mathrand.Rand
}

func (t testIntner) Intn(n int) (int, error) {
	return t.r.Intn(n), nil
}

var (
	testParts = []string{"correct", "horse", "battery", "staple"}
	testAlpha = []string{"!", "@", "£", "$"}
)
