package main

import "fmt"

func anotherNum() int {
	return 50
}

func main() {
	num := 42

	defer func() {
		fmt.Println("defer func: num:", num)

		fmt.Println("defer func : anotherNum():", anotherNum())
	}()

	num = 13

	fmt.Println("main: num:", num)
}
