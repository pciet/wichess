package main

import "fmt"

const (
	Debug    = true
	DebugSQL = false
)

func DebugPrintln(a ...interface{}) {
	if Debug {
		fmt.Println(a...)
	}
}
