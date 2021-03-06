// Coding Exercise #2

// Rename the file created at Exercise #1 to data.txt

// Check if file exists before renaming it. If it doesn't exist print a message and stop the program.

package main

import (
	"log"
	"os"
)

func main() {
	// checking if file exists
	_, err := os.Stat("info.txt")
	// error handling
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist!")
		}
	}

	// renaming the file
	err = os.Rename("info.txt", "data.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	newFile, err := os.Create("info.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	newFile.Close()

// 	oldFile := "info.txt"
// 	renamedFile := "data.txt"
// 	if renamedFile == oldFile {
// 		fmt.Println("File already exists!")
// 	} else {
// 		err = os.Rename(oldFile, renamedFile)
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
