package main

import (
	"fmt"
	"reflect" // for tags
)

type Doctor struct {
	number     int
	actorName  string
	companions []string //slice, not array
}

type Animal struct {
	Name   string `required max : "100"` // tag : required + max length
	Origin string
}

type Bird struct {
	Animal   // composition : has a : relationship
	SpeedKPH float32
	CanFly   bool
}

func main() {
	statePopulation := map[string]int{
		"Cali":     1000,
		"Texas":    90,
		"Flo":      811,
		"New York": 192,
	}

	m := make(map[string]int, 10)

	// add
	statePopulation["Geo"] = 10092018
	// order of keys will not be maintained
	fmt.Println(statePopulation, m)

	// del
	delete(statePopulation, "Geo")
	fmt.Println(statePopulation, m)

	fmt.Println(statePopulation["Geo"])

	// check if presence
	pop, ok := statePopulation["Flo"]
	fmt.Println(pop, ok)

	// STRUCTS
	aDoctor := Doctor{
		number:    3,
		actorName: "John",
		companions: []string{
			"Liz",
			"Jo",
			"Sarah",
		},
	}
	fmt.Println(aDoctor)

	// Way 1
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Aus"
	b.SpeedKPH = 60
	b.CanFly = false
	fmt.Println(b)

	// Way 2
	c := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Aus"},
		SpeedKPH: 60,
		CanFly:   false,
	}
	fmt.Println(c)

	// tagging - fectching tag
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
