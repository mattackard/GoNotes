//GoNotes
package main

import (
	"fmt"
	"os"

	"github.com/mattackard/project-0/notes"
)

func init() {
	loadConfig()
}

func main() {
	//Retrieves the cli cmd passed
	var cmd string
	if len(os.Args) < 2 {
		cmd = ""
	} else {
		cmd = os.Args[1]
	}

	//determines what function to run based on the cli cmds
	switch cmd {
	case "create":
		notes.CreateFile(cfg, os.Args[2], os.Args[3])
	case "config":
		notes.Config()
	case "edit":
		notes.Edit(os.Args[2])
	case "delete":
		notes.Delete(os.Args[2])
	default:
		if cmd == "" {
			fmt.Printf("You must enter a command. \n")
		} else {
			fmt.Printf("%s is not recognized as a command \n", cmd)
		}
		notes.Print("help.txt")
	}
}
