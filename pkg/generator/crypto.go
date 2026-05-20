package generator

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type cryptoRand struct{}

func (cryptoRand) Intn(n int) (int, error) {
	v, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, fmt.Errorf("crypto/rand: %w", err)
	}

	return int(v.Int64()), nil
}
