// Coding Exercise #1

// There is an error in the following Go Program. Even though the goroutine is correctly launched, it doesn't print any message.

//  package main
 
// import (
//     "fmt"
// )
 
// func sayHello(n string) {
//     fmt.Printf("Hello, %s!\n", n)
// }
 
// func main() {
//     go sayHello("Mr. Wick")
// }
// Your task is to synchronize main and the goroutine using WaitGroups. The program should print the string received as argument by sayHello().

package main

import (
	"fmt"
	"sync"
)

func sayHello(n string, w *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	w.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go sayHello("Mr. Wick", &wg)

	wg.Wait()
}
