//Package notes is a package to create, edit, and delete notes from the command line
package notes

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

//CreateFile creates a text file in the project directory
func CreateFile(config *ini.File, fileName string, text string) {

	//the underscore is an unused error variable return from Create()
	fileExtension := config.Section("options").Key("fileExtension").String()
	f, err := os.Create(fileName + fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	//writes a string to the file using the reference created with Create()
	f.WriteString(text)
}

//Config currently prints out the contents of your config file
//Eventually config will open up an editor for changing config settings
func Config() {
	Print("config.ini")
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
