//Package notes is a package to create, edit, and delete notes from the command line
package notes

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/ini.v1"
)

//CreateFile creates a text file in the project directory
func CreateFile(config *ini.File, fileName string, text string) {

	//the underscore is an unused error variable return from Create()
	fileExtension := config.Section("options").Key("fileExtension").String()
	f, _ := os.Create(fileName + fileExtension)

	//writes a string to the file using the reference created with Create()
	f.WriteString(text)
}

//Config currently prints out the contents of your config file
//Eventually config will open up an editor for changing config settings
func Config() {
	ini, err := ioutil.ReadFile("config.ini")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(ini), "\n")
}
