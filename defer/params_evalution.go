package main

import "fmt"

func count(i int) (n int) {
	//Go runtime will save any passed params to the deferred func at the time of registering the defer
	// — not when it runs.
	defer func(i int) {
		n = n + i
	}(i)

	i = i * 2
	n = i

	return
}

func main() {
	fmt.Println(count(10))
}
