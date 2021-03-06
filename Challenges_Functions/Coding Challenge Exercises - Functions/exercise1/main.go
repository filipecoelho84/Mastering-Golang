// Coding Exercise #1

// Create a function called cube() that takes a parameter of type float64 and returns the cube of that parameter (the parameter to the power of 3).

package main

import "fmt"

func cube(n float64) float64 {
	return n * n * n
}

func main() {

	x := cube(3)
	fmt.Println(x)

}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func cube(a float64) float64 {
// 	return math.Pow(a, 3)
// }

// func main() {
// 	fmt.Println(cube(2))
// }
