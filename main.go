//GoNotes
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mattackard/project-0/config"
	"github.com/mattackard/project-0/notes"
	"github.com/zserge/webview"
)

var cfg config.Config
var cmd string
var fullPath string
var open *bool

func init() {
	cfg = config.LoadConfig()

	//Retrieves the cli cmd passed
	if len(os.Args) < 2 {
		cmd = ""
	} else {
		cmd = os.Args[1]
	}

	//Removes cmd from os.Args passed in to allow for parsing command-dependent flags
	os.Args = os.Args[1:]

	//get flags and args from the original terminal call
	open = flag.Bool("open", false, "open file for editing after creating")
	flag.Parse()
	os.Args = flag.Args()

	//sets a variable to the full file path passed in through args
	//if the command takes a filepath
	if len(os.Args) > 1 {
		fullPath = cfg.Paths.Notes + os.Args[0] + cfg.Options.FileExtension
	}
}

func main() {

	//determines what function to run based on the cli cmds
	switch cmd {
	case "create":
		notes.CreateFile(cfg, fullPath, open)
	case "config":
		notes.Config()
	case "edit":
		notes.Edit(fullPath)
	case "delete":
		notes.Delete(fullPath)
	case "gui":
		go webview.Open("GoNotes", "file:///home/ubuntu/go/src/github.com/mattackard/project-0/gui/gui.html", 600, 800, true)
		http.HandleFunc("/newNote", func(w http.ResponseWriter, r *http.Request) {
			//fmt.Println("new note was clicked")
			fmt.Fprint(w, "Changed")
		})
		http.HandleFunc("/deleteNote", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("delete note was clicked")
			io.WriteString(w, "Changed")
		})
		http.HandleFunc("/saveNote", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("save note was clicked")
			io.WriteString(w, "Changed")
		})
		http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("settings was clicked")
			io.WriteString(w, "Changed")
		})
		fmt.Println("Server is running at localhost", cfg.Options.Port)
		http.ListenAndServe(cfg.Options.Port, nil)
	default:
		if cmd == "" {
			fmt.Printf("You must enter a command. \n")
		} else {
			fmt.Printf("%s is not recognized as a command \n", cmd)
		}
		notes.Print("help.txt")
	}
}
