package main

import (
	"fmt"
)

const Debug = true

func DebugPrintln(a ...interface{}) {
	if Debug {
		_, err := fmt.Println(a...)
		if err != nil {
			panic(err.Error())
		}
	}
}
