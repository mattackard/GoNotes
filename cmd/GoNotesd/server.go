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
	"github.com/mattackard/project-1/pkg/dnsutil"
	"github.com/mattackard/project-1/pkg/logutil"
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

var logFile *os.File

func main() {

	//open log file
	logFile = logutil.OpenLogFile("./logs/")

	// set up server endpoints using middleware
	connectHanlder := http.HandlerFunc(connect)
	http.Handle("/connect", authorizeRequest(sendToLogger(connectHanlder)))

	newNoteHanlder := http.HandlerFunc(newNote)
	http.Handle("/newNote", authorizeRequest(sendToLogger(newNoteHanlder)))

	noteDirHanlder := http.HandlerFunc(noteDir)
	http.Handle("/dir", authorizeRequest(sendToLogger(noteDirHanlder)))

	getFileHanlder := http.HandlerFunc(getFile)
	http.Handle("/getFile", authorizeRequest(sendToLogger(getFileHanlder)))

	deleteNoteHanlder := http.HandlerFunc(deleteNote)
	http.Handle("/deleteNote", authorizeRequest(sendToLogger(deleteNoteHanlder)))

	saveNoteHanlder := http.HandlerFunc(saveNote)
	http.Handle("/saveNote", authorizeRequest(sendToLogger(saveNoteHanlder)))

	settingsHanlder := http.HandlerFunc(settings)
	http.Handle("/settings", authorizeRequest(sendToLogger(settingsHanlder)))

	// start server on the port specified in the config file
	myIP := dnsutil.Ping("dns:6060", "noteserver")
	noPort := dnsutil.TrimPort(myIP)
	logutil.SendLog("logger:6060", false, []string{"NoteServer started at " + noPort + config.Mycfg.Options.Port}, logFile, "NoteServer")
	log.Println(http.ListenAndServe(config.Mycfg.Options.Port, nil))
}

func sendToLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logutil.LogServerRequest(w, r, "logger:6060", logFile, "NoteServer")
		next.ServeHTTP(w, r)
	})
}

// set header to expect json and allow cors
func setHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	return w
}

//checks if the request has the proper auth token in its header
func authorizeRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Proxy-Authorization")
		encode := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PROXYAUTH")))
		encode = "Basic " + encode

		//if auth doesn't match, reject request
		if auth != encode {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Invalid Authorization")
			return
		}
		next.ServeHTTP(w, r)
	})
}

//return the status of the connection from the client
func connect(w http.ResponseWriter, r *http.Request) {
	w = setHeaders(w)
	w.Write([]byte("OK"))
}

// create a new note with datestamp and returns it as http response
func newNote(w http.ResponseWriter, r *http.Request) {

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
		logutil.SendLog("logger:6060", true, []string{err.Error()}, logFile, "NoteServer")
	}

	w = setHeaders(w)
	w.Write(js)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {

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

func saveNote(w http.ResponseWriter, r *http.Request) {

	// parse request to get files name and text content to save
	var requestNote note
	save, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logutil.SendLog("logger:6060", true, []string{err.Error()}, logFile, "NoteServer")
	}
	json.Unmarshal(save, &requestNote)
	if err != nil {
		logutil.SendLog("logger:6060", true, []string{err.Error()}, logFile, "NoteServer")
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

// sends the config file back to the client
func settings(w http.ResponseWriter, r *http.Request) {

	// load config.json and marshal into JSON
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		logutil.SendLog("logger:6060", true, []string{err.Error()}, logFile, "NoteServer")
	}
	response := note{
		Path:     "./",
		FileName: "config.json",
		Text:     string(file),
	}
	js, err := json.Marshal(response)
	if err != nil {
		logutil.SendLog("logger:6060", true, []string{err.Error()}, logFile, "NoteServer")
	}

	w = setHeaders(w)
	w.Write(js)
}

// gets all items in the requested directory
func noteDir(w http.ResponseWriter, r *http.Request) {

	// Unmarshal post body to get requested directory
	var newDir directory
	save, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logutil.SendLog("logger:6060", true, []string{err.Error()}, logFile, "NoteServer")
	}
	json.Unmarshal(save, &newDir)
	if err != nil {
		logutil.SendLog("logger:6060", true, []string{err.Error()}, logFile, "NoteServer")
	}

	// get the list of files in the requested directory and send in response
	files := notes.List(newDir.Root)
	d := directory{
		Root:  newDir.Root,
		Files: files,
	}
	js, err := json.Marshal(d)
	if err != nil {
		logutil.SendLog("logger:6060", true, []string{err.Error()}, logFile, "NoteServer")
	}

	w = setHeaders(w)
	w.Write(js)
}

// gets the text content of a single file and returns it in response
func getFile(w http.ResponseWriter, r *http.Request) {

	// Unmarshal post body to get filename
	var requestFile note
	save, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logutil.SendLog("logger:6060", true, []string{err.Error()}, logFile, "NoteServer")
	}
	json.Unmarshal(save, &requestFile)
	if err != nil {
		logutil.SendLog("logger:6060", true, []string{err.Error()}, logFile, "NoteServer")
	}

	// Read file and marshal data in JSON for response
	file, err := ioutil.ReadFile(requestFile.Path + requestFile.FileName)
	if err != nil {
		logutil.SendLog("logger:6060", true, []string{err.Error()}, logFile, "NoteServer")
	}
	response := note{
		Path:     requestFile.Path,
		FileName: requestFile.FileName,
		Text:     string(file),
	}
	js, err := json.Marshal(response)
	if err != nil {
		logutil.SendLog("logger:6060", true, []string{err.Error()}, logFile, "NoteServer")
	}

	w = setHeaders(w)
	w.Write(js)
}
