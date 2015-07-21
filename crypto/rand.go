package crypto

import (
	"crypto/rand"
	"math/big"
)

// RandBytes returns n bytes of random data.
func RandBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

// RandIntn returns a non-negative random integer in the range [0,n). It
// panics if n <= 0.
func RandIntn(n int) (int, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	return int(r.Int64()), err
}