package rules

import (
	"crypto/rand"
	"math"
	"math/big"
	prand "math/rand"
)

var randomSource = func() *prand.Rand {
	seed, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		Panic(err)
	}
	return prand.New(prand.NewSource(seed.Int64()))
}()
