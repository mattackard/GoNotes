//GoNotes
package main

import (
	"fmt"
	"os"
)

func main() {
	//Retrieves the file name passed in as the first argument
	var cmd string = os.Args[1]

	//check the value of 1st arg to determine function to run
	switch cmd {
	case "create":
		createFile(os.Args[2], os.Args[3])
	case "config":
		config()
	default:
		fmt.Printf("%s is not recognized as a command \n", os.Args[1])
	}

}

func createFile(fileName, text string) {
	//creates a text file in the project directory
	//the underscore is an unused error variable return from Create()
	f, _ := os.Create(fileName + ".txt")

	//writes a string to the file using the reference created with Create()
	f.WriteString(text)
}

func config() {
	fmt.Println("config called")
}
