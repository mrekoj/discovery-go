package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("later")
	}()

	fmt.Println("first")
}
