package main

import (
	"fmt"
)

func main() {
	// copy by value
	a := 42
	b := a
	fmt.Println(a, b)
	a = 10
	fmt.Println(a, b)

	// copy by ref
	var c int = 42
	var d *int = &c
	fmt.Println(c, *d)
	c = 10
	fmt.Println(c, *d)
	*d = 1e6
	fmt.Println(c, *d)

	// pointer arithmetic to be avoided
	// if inevitable, import "unsafe"

	// uninitialized pointers = nil

}
