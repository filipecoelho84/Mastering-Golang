// Coding Exercise #5

// Change the function from the previous exercise and use a `naked return`.

package main

import "fmt"

func sum(a ...int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func main() {

	s := sum(1, 2, 30)
	fmt.Println(s)

}

// package main

// import "fmt"

// func sum(a ...int) (s int) {
// 	for _, v := range a {
// 		s += v
// 	}
// 	return
// }

// func main() {

// 	s := sum(1, 2, 30)
// 	fmt.Println(s)

// }
