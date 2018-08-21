package main

import "fmt"

type Car struct {
	model string
}

func (c Car) PrintModel() {
	fmt.Printf("Value-receiver see the model as '%s'\n", c.model)
}

func (c *Car) PrintModelViaPointer() {
	fmt.Printf("Pointer-receiver see the  model as '%s'\n", c.model)
}

func withoutPointer() {
	c := Car{model: "Ford Focus"}
	defer c.PrintModel()
	c.model = "Honda City"
}

func withPointer() {
	c := Car{model: "Ford Focus"}
	defer c.PrintModelViaPointer()
	c.model = "Honda City"
}

func main() {
	withoutPointer()
	withPointer()
}
