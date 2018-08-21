package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Proto string
}

func newFromFile(path string) (c *Config, err error) {
	c = &Config{}

	file, err := os.Open(path)

	if err != nil {
		return
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(c)
	if err != nil {
		return
	}

	return
}

func main() {
	cfg, err := newFromFile("config.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", cfg)
}
