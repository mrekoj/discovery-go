//https://dave.cheney.net/2014/03/19/channel-axioms

package main

import "fmt"

func main() {
	var c = make(chan int, 100)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			// Channel close in first loop, so the remain can't send value to channel c
			close(c)
		}()
	}

	for i := range c {
		fmt.Println(i)
	}
}
