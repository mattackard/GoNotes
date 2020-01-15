package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mattackard/project-0/config"
)

type note struct {
	FileName string `json:"fileName"`
	Title    string `json:"title"`
	Text     string `json:"text"`
}

var currentNote note

func main() {
	http.HandleFunc("/newNote", newNote)
	http.HandleFunc("/deleteNote", deleteNote)
	http.HandleFunc("/saveNote", saveNote)
	http.HandleFunc("/settings", settings)
	fmt.Println("Server is running at localhost", config.Mycfg.Options.Port)
	http.ListenAndServe(config.Mycfg.Options.Port, nil)
}

func setHeaders(w http.ResponseWriter) http.ResponseWriter {
	//set header to expect json and allow cors
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return w
}

func newNote(w http.ResponseWriter, r *http.Request) {
	//create a new note struct and build a json object
	newNote := note{
		FileName: "myfile.txt",
		Title:    "New Note",
		Text:     "New note endpoint",
	}
	js, err := json.Marshal(newNote)
	if err != nil {
		panic(err)
	}

	w = setHeaders(w)
	w.Write(js)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	//create a new note struct and build a json object
	newNote := note{
		FileName: "myfile.txt",
		Title:    "New Note",
		Text:     "Delete note endpoint",
	}
	js, err := json.Marshal(newNote)
	if err != nil {
		panic(err)
	}

	w = setHeaders(w)
	w.Write(js)
}

func saveNote(w http.ResponseWriter, r *http.Request) {
	//create a new note struct and build a json object
	newNote := note{
		FileName: "myfile.txt",
		Title:    "New Note",
		Text:     "Save note endpoint",
	}
	js, err := json.Marshal(newNote)
	if err != nil {
		panic(err)
	}

	w = setHeaders(w)
	w.Write(js)
}

func settings(w http.ResponseWriter, r *http.Request) {
	//create a new note struct and build a json object
	newNote := note{
		FileName: "config.json",
		Title:    "Settings",
		Text:     "Settings endpoint",
	}
	js, err := json.Marshal(newNote)
	if err != nil {
		panic(err)
	}

	w = setHeaders(w)
	w.Write(js)
}
