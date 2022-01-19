package main

import (
	"fmt"
)

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sum(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return result
}

// we can have multiple returns from one function

func main() {
	sayMessage("Hello Go!")
	sayMessage("Hello")
	fmt.Println(sum(1, 1, 2, 3, 4, 5, 6, 7, 3, 21, 2, 4, 12, 32, 21, 4))
}
