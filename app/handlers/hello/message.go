package main

import (
	"fmt"
	"runtime"
)

type Language int
const (
	Thai = iota
	English
)

func message(l Language) string{
	return fmt.Sprintf("%s %s", greet(l), runtime.Version())
}

func greet(l Language) string {
	switch l {
	case Thai:
		return "โย่วๆๆๆ"
	case English:
		return "Hello World"
	default:
		return "No language - I guess English then? Hello!"
	}
}