package rules

import "fmt"

// Panic uses the fmt.Println formatting and calls panic with the resulting string.
func Panic(a ...interface{}) { panic(fmt.Sprintln(a...)) }
