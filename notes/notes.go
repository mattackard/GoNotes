//Package notes is a package to create, edit, and delete notes from the command line
package notes

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mattackard/project-0/config"
)

//CreateFile creates a text file in the project directory
func CreateFile(config config.Config, fileName string, text string) {
	f, err := os.Create(fileName + config.Options.FileExtension)
	if err != nil {
		log.Fatal(err)
	}

	//writes a string to the file using the reference created with Create()
	f.WriteString(text)
}

//Config currently prints out the contents of your config file
//Eventually config will open up an editor for changing config settings
func Config() {
	Print("config.json")
}

//Print opens an existing file and prints the contents into the terminal
func Print(fileName string) {
	//reads the whole file and stores as a string in note
	note, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(note), "\n")
}

//Edit allows for editing and saving notes
func Edit(fileName string) {
	//Terminal UI library?
	Print(fileName)
}

//Delete removes the given file
func Delete(fileName string) {
	os.Remove(fileName)
	fmt.Printf("%s has been deleted \n", fileName)
}
