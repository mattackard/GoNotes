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

//Note holds file information for files used in the editor
type note struct {
	Path     string `json:"path"`
	FileName string `json:"fileName"`
	Text     string `json:"text"`
}

type directory struct {
	Root  string   `json:"root"`
	Files []string `json:"files"`
}

func main() {
	//set up server endpoints
	http.HandleFunc("/newNote", newNote)
	http.HandleFunc("/dir", noteDir)
	http.HandleFunc("/getFile", getFile)
	http.HandleFunc("/deleteNote", deleteNote)
	http.HandleFunc("/saveNote", saveNote)
	http.HandleFunc("/settings", settings)

	//start server on the port specified in the config file
	fmt.Println("Server is running at localhost", config.Mycfg.Options.Port)
	http.ListenAndServe(config.Mycfg.Options.Port, nil)
}

//set header to expect json and allow cors
func setHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	return w
}

//create a new note with datestamp and returns it as http response
func newNote(w http.ResponseWriter, r *http.Request) {
	//get current date and format it
	currentTime := time.Now()
	prettyTime := currentTime.Format("Mon January _2, 2006")

	//add date to top of file and add some newlines for formatting
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
	//read the request to get the filename to delete
	var requestNote note
	delBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(delBody, &requestNote)

	//delete the file
	filePath := requestNote.Path + requestNote.FileName + config.Mycfg.Options.FileExtension
	notes.Delete(filePath)

	//send back success response
	w = setHeaders(w)
	w.Write([]byte("OK"))
}

func saveNote(w http.ResponseWriter, r *http.Request) {
	//parse request to get files name and text content to save
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
		filePath = requestNote.Path + requestNote.FileName
	} else {
		filePath = requestNote.Path + requestNote.FileName + config.Mycfg.Options.FileExtension
	}
	notes.Update(requestNote.Path, filePath, requestNote.Text)

	//If file updated was config.json reload the global variable
	if requestNote.FileName == "config.json" {
		config.Mycfg = config.LoadConfig()
	}

	w = setHeaders(w)
	w.Write([]byte("OK"))
}

//sends the config file back to the client
func settings(w http.ResponseWriter, r *http.Request) {
	//load config.json and marshal into JSON
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	response := note{
		Path:     "./",
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

//gets all items in the requested directory
func noteDir(w http.ResponseWriter, r *http.Request) {
	//Unmarshal post body to get requested directory
	var newDir directory
	save, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(save, &newDir)
	if err != nil {
		panic(err)
	}

	//get the list of files in the requested directory and send in response
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

//gets the text content of a single file and returns it in response
func getFile(w http.ResponseWriter, r *http.Request) {
	//Unmarshal post body to get filename
	var requestFile note
	save, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(save, &requestFile)
	if err != nil {
		panic(err)
	}

	//Read file and marshal data in JSON for response
	file, err := ioutil.ReadFile(requestFile.Path + requestFile.FileName)
	if err != nil {
		panic(err)
	}
	response := note{
		Path:     requestFile.Path,
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
