package auth

import (
	"log"
	"os"
	"runtime/debug"
)

// authRecover is deferred to recover and print additional debugging information for panics within
// package auth handlers. The print is followed by a call to os.Exit(1).
func authRecover(a ...interface{}) {
	pv := recover()
	if pv == nil {
		return
	}
	log.Println(pv)
	log.Println(a...)
	debug.PrintStack()
	os.Exit(1)
}
