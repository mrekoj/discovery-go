//http://guzalexander.com/2013/12/06/golang-channels-tutorial.html
package main

import (
	"fmt"
	"time"
)

func main() {
	count := 3

	c := make(chan string, 2) // capacity = 0

	go func() {
		for i := 0; i < count; i++ {
			println("send message")
			c <- fmt.Sprintf("message%d", i)
		}
		//close(c)
	}()

	// If capacity = 0 then when channel received first message
	// it's sleep 3 second then pop out first value
	// Then continue...
	// So channel without capacity like a queue with one element
	// It's wait for element pop out then can have another element to push in
	// ==========================================
	// If capacity = 2 then all 3 message are write into channel
	// and main program wait 3 second and can read from channel
	time.Sleep(3 * time.Second)

	// if you not close channel then for range will throw error

	for s := range c {
		println(s)
	}
	//for i := 0; i < count; i++ {
	//	println(<-c)
	//}
}
