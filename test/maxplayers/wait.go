package main

import (
	"math/rand"
	"time"
)

const (
	MinResponseTimeMilliseconds = 300
	MaxResponseTimeMilliseconds = 3000
)

func RandomHumanWait() {
	msecs := rand.Intn(MaxResponseTimeMilliseconds - MinResponseTimeMilliseconds)
	<-time.After(time.Duration(MinResponseTimeMilliseconds+msecs) * time.Millisecond)
}
