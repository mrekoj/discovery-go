package main

import "fmt"

func stackDefer() {
	defer func() {
		fmt.Println("Second")
	}()

	defer func() {
		fmt.Println("First")
	}()
}

func main() {
	stackDefer()
}
