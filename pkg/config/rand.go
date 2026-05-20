package config

// Intner provides a random integer in [0, n).
type Intner interface {
	Intn(n int) (int, error)
}
