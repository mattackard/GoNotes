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

	//determines what function to run based on the cli cmds
	switch cmd {
	case "create":
		notes.CreateFile(cfg, os.Args[2], os.Args[3])
	case "config":
		notes.Config()
	case "edit":
		notes.Edit(cfg.Paths.Notes + os.Args[2] + cfg.Options.FileExtension)
	case "delete":
		notes.Delete(cfg.Paths.Notes + os.Args[2] + cfg.Options.FileExtension)
	default:
		if cmd == "" {
			fmt.Printf("You must enter a command. \n")
		} else {
			fmt.Printf("%s is not recognized as a command \n", cmd)
		}
		notes.Print("help.txt")
	}
}
