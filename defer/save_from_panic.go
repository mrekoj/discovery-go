package main

import (
	"fmt"
	"runtime/debug"
)

func saveMe() {
	defer func() {
		fmt.Printf("recovered in defer: \"%s\"\n\n", recover())
		debug.PrintStack()
	}()

	panic("oops!")
}

func main() {
	saveMe()
}
