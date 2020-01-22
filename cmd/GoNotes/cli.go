// GoNotes
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mattackard/project-0/pkg/config"
	"github.com/mattackard/project-0/pkg/notes"
)

// fullPath holds the path including directory to use when referencing note files
var fullPath string

// cmd holds the subcommand to run
var cmd string

func main() {
	// Retrieves the cli cmd passed
	if len(os.Args) < 2 {
		cmd = ""
	} else {
		cmd = os.Args[1]
	}

	// get flags and args from the original terminal call
	open := flag.Bool("open", config.Mycfg.Options.InitEditor, "open file for editing after creating")
	dateStamp := flag.Bool("date", config.Mycfg.Options.DateStamp, "Initializes new note files with the current date")
	if len(os.Args) > 2 {
		flag.CommandLine.Parse(os.Args[2:])
		os.Args = flag.Args()
	}

	// sets a variable to the full file path passed in through args
	// if the command takes a filepath
	if len(os.Args) > 0 {
		if strings.Contains(os.Args[0], ".") {
			fullPath = config.Mycfg.Paths.Notes + os.Args[0]
		} else {
			fullPath = config.Mycfg.Paths.Notes + os.Args[0] + config.Mycfg.Options.FileExtension
		}
	}

	// determines what function to run based on the cli cmds
	switch cmd {
	case "create":
		notes.CreateFile(config.Mycfg.Paths.Notes, fullPath, *open, *dateStamp)
	case "config":
		notes.Config()
	case "list":
		myDir := notes.List(config.Mycfg.Paths.Notes)
		for _, v := range myDir {
			fmt.Println(v)
		}
	case "edit":
		notes.Edit(fullPath)
	case "delete":
		notes.Delete(fullPath)
	default:
		// prints out a help message with the possible commands if an unrecognized command is entered
		if cmd == "" {
			fmt.Printf("You must enter a command. \n")
		} else {
			fmt.Printf("%s is not recognized as a command \n", cmd)
		}
		notes.Print("help.txt")
	}
}
