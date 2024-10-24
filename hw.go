package main

import (
	"fmt"
	"math"
)

var Global int = 1234
var AnotherGlobal = -5678 // Inferred based on the output type. We do not have to define as integer

func main() {
	var j int
	fmt.Println("Initial j value:", j) // Printing an uninitialized variable

	i := Global + AnotherGlobal
	j = Global
	k := math.Abs(float64(AnotherGlobal))

	// Print all the values after being initialized
	fmt.Printf("Global=%d, i=%d, j=%d, k=%.2f.\n", Global, i, j, k)

}
