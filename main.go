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
	var cmd string = os.Args[1]

	//determines what function to run based on the cli cmds
	switch cmd {
	case "create":
		notes.CreateFile(cfg, os.Args[2], os.Args[3])
	case "config":
		notes.Config()
	default:
		fmt.Printf("%s is not recognized as a command \n", cmd)
	}
}
