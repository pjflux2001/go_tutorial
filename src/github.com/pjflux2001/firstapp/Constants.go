package main

import (
	"fmt"
)

const a int16 = 17

func main() {
	const a int32 = 9 * iota
	fmt.Println(a)
}
