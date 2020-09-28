package game

import "fmt"

const debugEnabled = true

func debug(a ...interface{}) {
	if debugEnabled {
		fmt.Println(a...)
	}
}
