package main

func main() {
	done := make(chan bool)
	go func() {
		println("Goroutine done")
		//done <- true
		// Can you close(done) to notify channel not block
		close(done)
	}()

	println("Main done")
	<-done
}
