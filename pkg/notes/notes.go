//Package notes is a package to create, edit, and delete notes from the command line
package notes

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/mattackard/project-0/pkg/config"
)

//CreateFile creates a text file in the project directory
func CreateFile(config config.Config, filePath string, open *bool) {
	os.MkdirAll(config.Paths.Notes, 0777)

	//create the note file using the extension and path from config
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString(os.Args[2])

	//opens the file in the editor if open flag is set
	if *open {
		Edit(filePath)
	}
}

//Config opens the user's config file in the text editor
//It will also create a config file if it can't be found
func Config() {
	Edit("config.json")
}

//Print opens an existing file and prints the contents into the terminal
func Print(fileName string) string {
	//reads the whole file and stores as a byte[] in note
	note, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	//casts the byte[] to string for printing
	fmt.Print(string(note), "\n")
	return string(note)
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
		log.Fatal(err)
	}
}

//Delete removes the given file
func Delete(fileName string) {
	os.Remove(fileName)
}

//Update overwrites the given file with the new content
func Update(config config.Config, fileName string, text string) {
	os.MkdirAll(config.Paths.Notes, 0777)
	ioutil.WriteFile(fileName, []byte(text), 0777)
}

//List returns a slice of string file names of all files in the given directory
func List(directory string) []string {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	for _, v := range files {
		fileNames = append(fileNames, v.Name())
	}
	return fileNames
}
