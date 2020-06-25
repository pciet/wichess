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

func RandomSpecialKind() Kind {
	return Kind(randomSource.Int63n(int64(KindCount-BasicKindCount-1)) + 1 + BasicKindCount)
}

func TwoDifferentSpecialKinds() (Kind, Kind) {
	pa := RandomSpecialKind()
	pb := RandomSpecialKind()
	if pa == pb {
		pb++
		if pb == KindCount {
			pb = BasicKindCount + 1
		}
	}
	return pa, pb
}

func (a Kind) DifferentSpecialKind() Kind {
	p := RandomSpecialKind()
	if p == a {
		p++
		if p == KindCount {
			p = BasicKindCount + 1
		}
	}
	return p
}
