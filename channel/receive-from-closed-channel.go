// https://dave.cheney.net/2014/03/19/channel-axioms
package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	// When channel is closed, and all values drains from its buffer
	// The channel always return 0
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", <-c)
	}

	// Right solution
	//for i := range c {
	//	fmt.Println(i)
	//}
}
