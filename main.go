//GoNotes
package main

import (
	"os"
)

func main() {
	//Retrieves the file name passed in as the first argument
	var fileName string = os.Args[1]

	//creates a text file in the project directory
	//the underscore is an unused error variable return from Create()
	f, _ := os.Create(fileName + ".txt")

	//writes a string to the file using the reference created with Create()
	f.WriteString(os.Args[2])
}
