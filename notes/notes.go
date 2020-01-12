//Package notes is a package to create, edit, and delete notes from the command line
package notes

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/mattackard/project-0/config"
)

//CreateFile creates a text file in the project directory
func CreateFile(config config.Config, fileName string, text string) {
	os.MkdirAll(config.Paths.Notes, 0777)
	f, err := os.Create(config.Paths.Notes + fileName + config.Options.FileExtension)
	if err != nil {
		log.Fatal(err)
	}

	//writes a string to the file using the reference created with Create()
	f.WriteString(text)
}

//Config opens the user's config file in the text editor
//It will also create a config file if it can't be found
func Config() {
	Edit("config.json")
}

//Print opens an existing file and prints the contents into the terminal
func Print(fileName string) {
	//reads the whole file and stores as a byte[] in note
	note, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	//casts the byte[] to string for printing
	fmt.Print(string(note), "\n")
}

//Edit allows for editing and saving notes
func Edit(fileName string) {
	cmd := exec.Command("nano", fileName)

	//Exec defaults Stdin, out, err to dev/null unless specified
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//Open file in nano
	err := cmd.Run()
	if err != nil {
		println("Eror is not nil")
		log.Fatal(err)
	}
}

//Delete removes the given file
func Delete(fileName string) {
	os.Remove(fileName)
	fmt.Printf("%s has been deleted \n", fileName)
}
