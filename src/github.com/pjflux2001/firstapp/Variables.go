// all variable must be used : unused variables are to be removed
package main

import "fmt"

var i float64 = 100 // shadowed by local variable i

var I int = 420 // global exportable variable //

// 3 possible scope:
// - block
// - lowercase : scoped to package
// - uppercase
// NO private scope

var (
	actorName string = "SRK"
	companion string = "KS"
	number    int    = 10
)

func main() {
	// decl 1
	var i int
	i = 42
	fmt.Println(i)

	// decl 2
	var j float32 = 50
	j = float32(i)
	fmt.Println(j)

	// decl 3
	p := 69.0
	fmt.Println(p)

	// %value, %Type
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", p, p)
}
