package main

import (
	"fmt"
)

func main() {
	grades := [...]int{97, 99, 93}
	fmt.Println(grades)

	var students [3]string
	fmt.Println(students)

	students[0] = "S1"
	fmt.Println(students)
	fmt.Println(students[0])
	fmt.Println(len(students))

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	// deep copy by default
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 69
	fmt.Println(a, b)

	// for soft copy, use pointers
	c := &a
	c[1] = 1000000
	fmt.Println(a, b, c)

	// array : [...]
	// slice : []
	// most things about the two are same

	// slicing : slices and arrays works with both
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s[:])
	fmt.Println(s[3:])
	fmt.Println(s[:6])
	fmt.Println(s[3:6])
	// [inclusive : exclusive]

	// builtin make() : makes slice
	t := make([]int, 3)
	fmt.Println(t, len(t), cap(t))

	// **arrays have fixed size whereas slices have variable dynamic length**
	t = append(t, 1, 100, 1092, 1084, 994, 324)
	// append can't accept slice as a parameter but we can use spread operator
	t = append(t, []int{1, 100, 1092, 1084, 994, 324}...)
	fmt.Println(t, len(t), cap(t))
	t[5] = 7
	fmt.Println(t, len(t), cap(t))

	// remove from begin
	// a := a[1:]
	// remove from end
	// a := a[:len(a)-1]
	// remove from somewhere else
	// a := append(a[:2], a[3:]...)

}
