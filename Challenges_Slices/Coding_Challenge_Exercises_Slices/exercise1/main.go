// Coding Exercise #1

// Using a composite literal declare and initialize a slice of type string with 3 elements.

// Iterate over the slice and print each element in the slice and its index.

package main

import "fmt"

func main() {
	countries := []string{"Romania", "Brazil", "Germany"}
	for i, v := range countries {
		fmt.Printf("Index: %d, Element: %q\n", i, v)
	}

}

// package main

// import "fmt"

// func main() {
// 	slice := []string{"Filipe", "Sara", "Alice"}

// 	for index, value := range slice {
// 		if index < len(slice) {
// 			fmt.Printf("Index: %v, Element: %v\n", index, value)
// 		}
// 	}
// }
