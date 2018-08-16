//http://guzalexander.com/2013/12/06/golang-channels-tutorial.html
package main

import (
	"fmt"
	"time"
)

func getMessagesChannel(msg string, delay time.Duration) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; i <= 3; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Millisecond * delay)
		}
	}()

	return c
}

func main() {
	c1 := getMessagesChannel("first", 300)
	c2 := getMessagesChannel("second", 150)
	c3 := getMessagesChannel("third", 10)

	// Select will communication between all channel, it's print the first message come to channel
	for i := 1; i <= 9; i++ {
		select {
		case msg := <-c1:
			println(msg)
		case msg := <-c2:
			println(msg)
		case msg := <-c3:
			println(msg)

		}
	}
	//for i := 1; i <= 3; i++ {
	//	println(<-c1)
	//	println(<-c2)
	//	println(<-c3)
	//}
}
