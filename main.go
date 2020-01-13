//GoNotes
package main

import (
	"fmt"
	"os"

	"github.com/mattackard/project-0/config"
	"github.com/mattackard/project-0/notes"
)

var cfg config.Config

func init() {
	cfg = config.LoadConfig()
}

func main() {
	//Retrieves the cli cmd passed
	var cmd string
	if len(os.Args) < 2 {
		cmd = ""
	} else {
		cmd = os.Args[1]
	}

	//Removes cmd from os.Args passed in to allow for parsing command-dependent flags
	os.Args = os.Args[1:]

	//sets a variable to the full file path passed in through args
	//if the command takes a filepath
	var fullPath string
	if len(os.Args) > 1 {
		fullPath = cfg.Paths.Notes + os.Args[1] + cfg.Options.FileExtension
	}

	//determines what function to run based on the cli cmds
	switch cmd {
	case "create":
		notes.CreateFile(cfg, fullPath)
	case "config":
		notes.Config()
	case "edit":
		notes.Edit(fullPath)
	case "delete":
		notes.Delete(fullPath)
	default:
		if cmd == "" {
			fmt.Printf("You must enter a command. \n")
		} else {
			fmt.Printf("%s is not recognized as a command \n", cmd)
		}
		notes.Print("help.txt")
	}
}
