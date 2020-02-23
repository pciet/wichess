package main

import "log"

const (
	Debug    = true
	DebugSQL = false
)

func DebugPrintln(a ...interface{}) {
	if Debug {
		log.Println(a...)
	}
}
