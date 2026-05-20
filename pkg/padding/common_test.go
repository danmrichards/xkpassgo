package padding

import (
	mathrand "math/rand"
	"sync"
)

type syncIntner struct {
	mu sync.Mutex
	r  *mathrand.Rand
}

func (s *syncIntner) Intn(n int) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.r.Intn(n), nil
}

var (
	testParts = []string{"correct", "horse", "battery", "staple"}
	testAlpha = []string{"!", "@", "£", "$"}
)
