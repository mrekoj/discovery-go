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

	// main function is single process, so when write to channel c
	// main process are block util another process read data from channel c
	// so -> program block forever -> deadlock
	var c1 chan string
	fmt.Println(<-c1)

	//c:= make(chan string)
	//go func() {
	//	c <- "channel can send string value to it"
	//}()
	//
	//fmt.Println(<-c)
}
