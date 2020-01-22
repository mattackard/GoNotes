// Package notes is a package to create, edit, and delete files as well as print the contents of a file and output all files in a directory
package notes

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

// CreateFile creates a text file in directory defined in the user config
func CreateFile(path string, filePath string, open bool, date bool) {
	textContent := ""
	os.MkdirAll(path, 0777)

	// create the note file using the extension and path from config
	f, err := os.Create(filePath)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	// adds a date header to the file if the date flag is set
	if date {
		// get current date and format it
		currentTime := time.Now()
		prettyTime := currentTime.Format("Mon January _2, 2006")
		textContent = prettyTime + ",\n\n"
	}
	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			textContent += v + " "
		}
	}

	f.WriteString(textContent)

	// opens the file in the editor if open flag is set
	if open {
		Edit(filePath)
	}
}

// Config opens the user's config file in the text editor
// If no config file can be found a default will be created and then opened
func Config() {
	Edit("config.json")
}

// Print opens an existing file and prints the contents into the terminal
func Print(fileName string) string {
	// reads the file and stores as a byte[] in note
	note, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// casts the byte[] to string for printing
	fmt.Print(string(note), "\n")
	return string(note)
}

// Edit opens the given file in nano for editing and saving
func Edit(fileName string) {
	cmd := exec.Command("nano", fileName)

	// Exec defaults Stdin, out, err to dev/null unless specified
	// so you need to explicitly set the io
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Open file in nano
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// Delete removes the given file
func Delete(fileName string) {
	os.Remove(fileName)
}

// Update overwrites the given file with new text content
// can also be used to create a file
func Update(path string, fileName string, text string) {
	os.MkdirAll(path, 0777)
	err := ioutil.WriteFile(fileName, []byte(text), 0777)
	if err != nil {
		panic(err)
	}
}

// List returns a slice of string file names of all files in the given directory
func List(directory string) []string {
	// get all item in passed in directory
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	// put all file structs into a simple slice of filenames
	var fileNames []string
	for _, v := range files {
		fileNames = append(fileNames, v.Name())
	}
	return fileNames
}
