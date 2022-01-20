package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello!")
}

// interface defines behaviours
type Writer interface {
	Write([]_byte (int, error))
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}