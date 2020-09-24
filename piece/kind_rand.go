package piece

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

// RandomSpecialKind returns a randomly picked kind that isn't from a normal chess army.
func RandomSpecialKind() Kind {
	return Kind(randomSource.Int63n(int64(KindCount-basicKindCount-1)) + 1 + basicKindCount)
}

// TwoDifferentSpecialKinds returns two random special kinds that aren't the same.
func TwoDifferentSpecialKinds() (Kind, Kind) {
	pa := RandomSpecialKind()
	pb := RandomSpecialKind()
	if pa == pb {
		pb++
		if pb == KindCount {
			pb = basicKindCount + 1
		}
	}
	return pa, pb
}

// DifferentSpecialKind returns a random kind that's not the same.
func (a Kind) DifferentSpecialKind() Kind {
	p := RandomSpecialKind()
	if p == a {
		p++
		if p == KindCount {
			p = basicKindCount + 1
		}
	}
	return p
}
