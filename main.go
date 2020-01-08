//GoNotes
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

var cfg *ini.File

func init() {
	//attempt to load the config file
	config, err := ini.Load("config.ini")
	cfg = config

	//config error handling
	if err != nil {
		switch err.(type) {

		//if config.ini can't be found, create it
		case *os.PathError:
			println("Found a path error! Creating your .ini file.")
			os.Create("config.ini")

		default:
			log.Fatal(err)
		}
	}
}

func main() {
	//Retrieves the cli cmd passed
	var cmd string = os.Args[1]

	//determines what function to run based on the cli cmds
	switch cmd {
	case "create":
		createFile(os.Args[2], os.Args[3])
	case "config":
		config()
	default:
		fmt.Printf("%s is not recognized as a command \n", cmd)
	}

}

//Creates a text file in the project directory
func createFile(fileName, text string) {

	//the underscore is an unused error variable return from Create()
	fileExtension := cfg.Section("options").Key("fileExtension").String()
	f, _ := os.Create(fileName + fileExtension)

	//writes a string to the file using the reference created with Create()
	f.WriteString(text)
}

//Config currently prints out the contents of your config file
//Eventually config will open up an editor for changing config settings
func config() {
	ini, err := ioutil.ReadFile("config.ini")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(ini), "\n")
}
