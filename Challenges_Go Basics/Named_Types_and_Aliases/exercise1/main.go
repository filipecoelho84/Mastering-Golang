// Coding Exercise #1

// Declare a new type type called duration. Have the underlying type be an int.

// Declare a variable of the new type called hour using the var keyword

// In function main:

// print out the value of the variable hour

// print out the type of the variable hour

// assign 3600 to the variable hour using the  = operator

// print out the value of hour

package main

import "fmt"

type duration int

func main() {
	var hour duration
	fmt.Printf("hour's type: %T, hour's value: %v\n", hour, hour)
	hour = 3600
	fmt.Printf("hour's value %v\n", hour)
}
