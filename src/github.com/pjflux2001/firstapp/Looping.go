package main

import (
	"fmt"
)

func main() {
	for i, j := 0, 1; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	statePopulation := map[string]int{
		"Cali":     1000,
		"Texas":    90,
		"Flo":      811,
		"New York": 192,
	}
	for k, v := range statePopulation {
		fmt.Println(k, v)
	}

	for k, v := range "GeeksForGeeks" {
		fmt.Println(k, v)
	}

	for _, v := range "GeeksForGeeks" {
		fmt.Println(v)
	}
}
