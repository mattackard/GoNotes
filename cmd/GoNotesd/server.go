package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"time"

	"github.com/mattackard/project-0/pkg/config"
	"github.com/mattackard/project-0/pkg/notes"
)

// Note holds file information for files used in the editor
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

	// set up server endpoints
	http.HandleFunc("/connect", connect)
	http.HandleFunc("/newNote", newNote)
	http.HandleFunc("/dir", noteDir)
	http.HandleFunc("/getFile", getFile)
	http.HandleFunc("/deleteNote", deleteNote)
	http.HandleFunc("/saveNote", saveNote)
	http.HandleFunc("/settings", settings)

	// start server on the port specified in the config file
	fmt.Println("Server is running at http://server", config.Mycfg.Options.Port)
	log.Println(http.ListenAndServe(config.Mycfg.Options.Port, nil))
}

// set header to expect json and allow cors
func setHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	return w
}

//checks if the request has the proper auth token in its header
func authorizeRequest(w http.ResponseWriter, r *http.Request) bool {
	auth := r.Header.Get("Proxy-Authorization")
	encode := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PROXYAUTH")))
	encode = "Basic " + encode

	//if auth doesn't match, reject request
	if auth != encode {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid Authorization")
		return false
	}
	return true
}

//return the status of the connection from the client
func connect(w http.ResponseWriter, r *http.Request) {
	if authorizeRequest(w, r) {
		method := r.Method
		url := r.URL
		httpVer := r.Proto
		host := r.Host
		closeConn := r.Close
		address := r.RemoteAddr
		fmt.Println(method, url, httpVer, host, closeConn, address)
		w = setHeaders(w)
		w.Write([]byte("OK"))
	}
}

// create a new note with datestamp and returns it as http response
func newNote(w http.ResponseWriter, r *http.Request) {
	if authorizeRequest(w, r) {

		// get current date and format it
		currentTime := time.Now()
		prettyTime := currentTime.Format("Mon January _2, 2006")

		// add date to top of file and add some newlines for formatting
		os.MkdirAll(config.Mycfg.Paths.Notes, 0777)
		response := note{
			Path:     config.Mycfg.Paths.Notes,
			FileName: "",
			Text:     prettyTime + ", \n\n",
		}
		js, err := json.Marshal(response)
		if err != nil {
			log.Fatalln(err)
		}

		w = setHeaders(w)
		w.Write(js)
	}
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	if authorizeRequest(w, r) {

		// read the request to get the filename to delete
		var requestNote note
		delBody, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal(delBody, &requestNote)

		// delete the file
		var filePath string
		if strings.Contains(requestNote.FileName, ".") {
			filePath = requestNote.Path + requestNote.FileName
		} else {
			filePath = requestNote.Path + requestNote.FileName + config.Mycfg.Options.FileExtension
		}
		notes.Delete(filePath)

		// send back success response
		w = setHeaders(w)
		w.Write([]byte("OK"))
	}
}

func saveNote(w http.ResponseWriter, r *http.Request) {
	if authorizeRequest(w, r) {

		// parse request to get files name and text content to save
		var requestNote note
		save, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal(save, &requestNote)
		if err != nil {
			log.Fatalln(err)
		}

		// If a file extension is entered, use it. Otherwise use the extension from config
		// Keeps the config.json file in the project root
		var filePath string
		if requestNote.FileName == "config.json" {
			filePath = "./config.json"
		} else if strings.Contains(requestNote.FileName, ".") {
			filePath = requestNote.Path + requestNote.FileName
		} else {
			filePath = requestNote.Path + requestNote.FileName + config.Mycfg.Options.FileExtension
		}
		notes.Update(requestNote.Path, filePath, requestNote.Text)

		// If file updated was config.json reload the global variable
		if requestNote.FileName == "config.json" {
			config.Mycfg = config.LoadConfig()
		}

		w = setHeaders(w)
		w.Write([]byte("OK"))
	}
}

// sends the config file back to the client
func settings(w http.ResponseWriter, r *http.Request) {
	if authorizeRequest(w, r) {

		// load config.json and marshal into JSON
		file, err := ioutil.ReadFile("config.json")
		if err != nil {
			log.Fatalln(err)
		}
		response := note{
			Path:     "./",
			FileName: "config.json",
			Text:     string(file),
		}
		js, err := json.Marshal(response)
		if err != nil {
			log.Fatalln(err)
		}

		w = setHeaders(w)
		w.Write(js)
	}
}

// gets all items in the requested directory
func noteDir(w http.ResponseWriter, r *http.Request) {
	if authorizeRequest(w, r) {

		// Unmarshal post body to get requested directory
		var newDir directory
		save, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal(save, &newDir)
		if err != nil {
			log.Fatalln(err)
		}

		// get the list of files in the requested directory and send in response
		files := notes.List(newDir.Root)
		d := directory{
			Root:  newDir.Root,
			Files: files,
		}
		js, err := json.Marshal(d)
		if err != nil {
			log.Fatalln(err)
		}

		w = setHeaders(w)
		w.Write(js)
	}
}

// gets the text content of a single file and returns it in response
func getFile(w http.ResponseWriter, r *http.Request) {
	if authorizeRequest(w, r) {

		// Unmarshal post body to get filename
		var requestFile note
		save, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal(save, &requestFile)
		if err != nil {
			log.Fatalln(err)
		}

		// Read file and marshal data in JSON for response
		file, err := ioutil.ReadFile(requestFile.Path + requestFile.FileName)
		if err != nil {
			log.Fatalln(err)
		}
		response := note{
			Path:     requestFile.Path,
			FileName: requestFile.FileName,
			Text:     string(file),
		}
		js, err := json.Marshal(response)
		if err != nil {
			log.Fatalln(err)
		}

		w = setHeaders(w)
		w.Write(js)
	}
}
