package wichess

import (
	"crypto/rand"
	"math"
	"math/big"
	prand "math/rand"
)

var randomSource = func() *prand.Rand {
	seed, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err.Error())
	}
	return prand.New(prand.NewSource(seed.Int64()))
}()

func randomBool() bool {
	return randomSource.Intn(2) == 0
}
