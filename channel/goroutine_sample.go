//http://guzalexander.com/2013/12/06/golang-channels-tutorial.html

package main

/*
Main function don't wait goroutine finish it's work

*/
func main() {
	go println("goroutine message")

	println("main function message")
}
