package main

import "fmt"

func DebugPrintln(a ...interface{}) {
	if Debug {
		fmt.Println(a...)
	}
}
