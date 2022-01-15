package main

import (
	"fmt"
)

func main() {
	if 3 > 5 {
		fmt.Println("The test is True.")
	} else {
		fmt.Println("False.")
	}

	statePopulation := map[string]int{
		"Cali":     1000,
		"Texas":    90,
		"Flo":      811,
		"New York": 192,
	}
	// init; condition
	if pop, ok := statePopulation["Flo"]; ok {
		fmt.Println(pop)
	}

	val := 69
	// break(s) aren't required
	switch val {
	case 69:
		fmt.Println("a")
	case 2:
		fmt.Println("b")
	default:
		fmt.Println("l")
	}

	// tagless switch
	i := 10
	switch {
	case i <= 10:
		fmt.Println("Less than eq 10")
		fallthrough // anti-break (break was implicit)
	case i <= 20:
		fmt.Println("Less than eq 20")
	}
}
