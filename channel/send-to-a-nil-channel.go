//https://dave.cheney.net/2014/03/19/channel-axioms

package main

import "fmt"

func main() {
	/*
		If channel is unbuffered then a send to channel
		will block until another goroutine is ready to receive
	*/

	var c chan string

	c <- "let's get started"

	// The same with case receive from a nil channel block forever
	// until another goroutine send data to it
	var c1 chan string
	fmt.Println(<-c1)

	//c:= make(chan string)
	//go func() {
	//	c <- "channel can send string value to it"
	//}()
	//
	//fmt.Println(<-c)
}
