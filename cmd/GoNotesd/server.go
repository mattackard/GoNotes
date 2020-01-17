package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/mattackard/project-0/pkg/config"
	"github.com/mattackard/project-0/pkg/notes"
)

type note struct {
	FileName string `json:"fileName"`
	Text     string `json:"text"`
}

type directory struct {
	Root  string   `json:"root"`
	Files []string `json:"files"`
}

func main() {
	http.HandleFunc("/newNote", newNote)
	http.HandleFunc("/dir", noteDir)
	http.HandleFunc("/getFile", getFile)
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

	filePath := config.Mycfg.Paths.Notes + requestNote.FileName + config.Mycfg.Options.FileExtension
	notes.Delete(filePath)

	w = setHeaders(w)
	w.Write([]byte("OK"))
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

	//If a file extension is entered, use it. Otherwise use the extension from config
	//Keeps the config.json file in the project root
	var filePath string
	if requestNote.FileName == "config.json" {
		filePath = "./config.json"
	} else if strings.Contains(requestNote.FileName, ".") {
		filePath = config.Mycfg.Paths.Notes + requestNote.FileName
	} else {
		filePath = config.Mycfg.Paths.Notes + requestNote.FileName + config.Mycfg.Options.FileExtension
	}
	notes.Update(config.Mycfg, filePath, requestNote.Text)

	//If file updated was config.json reload the global variable
	if requestNote.FileName == "config.json" {
		config.Mycfg = config.LoadConfig()
	}

	w = setHeaders(w)
	w.Write([]byte("OK"))
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

func noteDir(w http.ResponseWriter, r *http.Request) {
	//Unmarshal post body
	var newDir directory
	save, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(save, &newDir)
	if err != nil {
		panic(err)
	}

	files := notes.List(newDir.Root)
	d := directory{
		Root:  newDir.Root,
		Files: files,
	}
	js, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	w = setHeaders(w)
	w.Write(js)
}

func getFile(w http.ResponseWriter, r *http.Request) {
	//Unmarshal post body
	var requestFile note
	save, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(save, &requestFile)
	if err != nil {
		panic(err)
	}

	//Create response and marshal to json
	file, err := ioutil.ReadFile(requestFile.FileName)
	if err != nil {
		panic(err)
	}
	response := note{
		FileName: requestFile.FileName,
		Text:     string(file),
	}
	js, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w = setHeaders(w)
	w.Write(js)
}
