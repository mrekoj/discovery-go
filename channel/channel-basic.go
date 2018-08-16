//http://guzalexander.com/2013/12/06/golang-channels-tutorial.html
package main

import "os"

func main() {
	i := make(chan int) // capacity = 0 by default

	s := make(chan string, 3) // non-zero capacity

	r := make(<-chan bool) // Only read from this channel

	w := make(chan<- os.FileInfo) // Only write to this channel

	readFromChannel(r)
	readFromChannel(w) // Accept only read channel

	// A channel which
	// - Only write to it
	// - It's value is another channel
	c := make(chan<- chan bool)

}

func readFromChannel(input <-chan bool) {}

func getChannel() chan bool {
	b := make(chan bool)
	return b
}
