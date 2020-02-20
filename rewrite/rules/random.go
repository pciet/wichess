package rules

import (
	"crypto/rand"
	"log"
	"math"
	"math/big"
	prand "math/rand"
)

var randomSource = func() *prand.Rand {
	seed, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		log.Panic(err)
	}
	return prand.New(prand.NewSource(seed.Int64()))
}()
