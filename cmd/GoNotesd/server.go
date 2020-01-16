package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mattackard/project-0/pkg/config"
	"github.com/mattackard/project-0/pkg/notes"
)

type note struct {
	FileName string `json:"fileName"`
	Text     string `json:"text"`
}

var noteDir string = config.Mycfg.Paths.Notes
var extension string = config.Mycfg.Options.FileExtension

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
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	return w
}

func newNote(w http.ResponseWriter, r *http.Request) {
	//create a new note struct and build a json object
	currentTime := time.Now()
	prettyTime := currentTime.Format("Mon January _2, 2006")
	response := note{
		FileName: "",
		Text:     prettyTime + ", \n\n",
	}
	js, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w = setHeaders(w)
	w.Write(js)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	//create a new note struct and build a json object
	var requestNote note
	delBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(delBody, &requestNote)

	filePath := noteDir + requestNote.FileName + extension
	notes.Delete(filePath)

	fmt.Fprint(w, "OK")
}

func saveNote(w http.ResponseWriter, r *http.Request) {
	//create a new note struct and build a json object
	var requestNote note
	save, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(save, &requestNote)
	if err != nil {
		panic(err)
	}

	filePath := noteDir + requestNote.FileName + extension
	notes.Update(filePath, requestNote.Text)

	fmt.Fprint(w, "OK")
}

func settings(w http.ResponseWriter, r *http.Request) {
	//create a new note struct and build a json object
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	response := note{
		FileName: "config.json",
		Text:     string(file),
	}
	js, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w = setHeaders(w)
	w.Write(js)
}
