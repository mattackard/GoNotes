package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/mattackard/project-0/pkg/config"
)

func TestSetHeaders(t *testing.T) {
	w := httptest.NewRecorder()
	setHeaders(w)
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content type set to (%s), expecting application/json", w.Header().Get("Content-Type"))
	}
	if w.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("Access-Control-Allow-Origin set to (%s), expecting (*)", w.Header().Get("Access-Control-Allow-Origin"))
	}
	if w.Header().Get("Access-Control-Allow-Methods") != "GET, POST" {
		t.Errorf("Access-Control-Allow-Methods set (%s), expecting (GET, POST)", w.Header().Get("Access-Control-Allow-Methods"))
	}
}

func TestNewNote(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	newNote(w, r)
	if w.Code != 200 {
		t.Errorf("Got error code %d, expecting 200", w.Code)
	}
	bytes, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Error("Error reading test response body")
	}
	var testBody note
	json.Unmarshal(bytes, &testBody)
	if testBody.FileName != "" {
		t.Errorf("Response filename is (%s), expecting blank string", testBody.FileName)
	}
	if testBody.Path != config.Mycfg.Paths.Notes {
		t.Errorf("Response path is (%s), expecting (%s)", testBody.Path, config.Mycfg.Paths.Notes)
	}
	if testBody.Text == "" {
		t.Error("Response text is blank")
	}
}

func TestDeleteNote(t *testing.T) {
	os.Create("test.txt")
	rBody := strings.NewReader(`{
		"path": "./",
		"fileName": "test",
		"text": ""
	  }`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", rBody)
	deleteNote(w, r)
	if w.Code != 200 {
		t.Errorf("Got error code %d, expecting 200", w.Code)
	}
	_, err := os.Open("test.txt")
	if err == nil {
		t.Error("The test file was not deleted")
	}
}

func TestSaveNote(t *testing.T) {
	rBody := strings.NewReader(`{
		"path": "./",
		"fileName": "test.md",
		"text": "test text"
	  }`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", rBody)
	saveNote(w, r)
	if w.Code != 200 {
		t.Errorf("Got error code %d, expecting 200", w.Code)
	}
	saved, err := os.Open("test.md")
	if err != nil {
		t.Error("Could not open the saved file")
	}
	text, err := ioutil.ReadAll(saved)
	if err != nil {
		t.Error("Could not read from saved file")
	}
	if string(text) != "test text" {
		t.Errorf("Saved test was (%s), expected (test text)", string(text))
	}
	os.Remove("test.md")
}

func TestSettings(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	settings(w, r)
	if w.Code != 200 {
		t.Errorf("Got error code %d, expecting 200", w.Code)
	}
	bytes, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Error("Error reading test response body")
	}
	var testBody note
	json.Unmarshal(bytes, &testBody)
	if testBody.FileName != "config.json" {
		t.Errorf("Response filename is (%s), expecting (config.json)", testBody.FileName)
	}
	if testBody.Path != "./" {
		t.Errorf("Response path is (%s), expecting (./)", testBody.Path)
	}
	if testBody.Text == "" {
		t.Error("Response text is blank")
	}
}

func TestNoteDir(t *testing.T) {
	rBody := strings.NewReader(`{
		"root": "./",
		"files": ""
	  }`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", rBody)
	noteDir(w, r)
	if w.Code != 200 {
		t.Errorf("Got error code %d, expecting 200", w.Code)
	}
	bytes, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Error("Error reading test response body")
	}
	var testDir directory
	json.Unmarshal(bytes, &testDir)
	if testDir.Root != "./" {
		t.Errorf("Response root is (%s), expected (./)", testDir.Root)
	}
	//GoNotesd directory is expected to have: Dockerfile, server_test.go, server.go, generated config.json, /notes
	if len(testDir.Files) != 5 {
		t.Error("Test directory is not expected length")
	}
}

func TestGetFile(t *testing.T) {
	rBody := strings.NewReader(`{
		"path": "./",
		"fileName": "config.json",
		"text": ""
	  }`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", rBody)
	getFile(w, r)
	if w.Code != 200 {
		t.Errorf("Got error code %d, expecting 200", w.Code)
	}
	bytes, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Error("Error reading test response body")
	}
	var testBody note
	json.Unmarshal(bytes, &testBody)
	if testBody.Path != "./" {
		t.Errorf("Response path is (%s), expected (./)", testBody.Path)
	}
	if testBody.FileName != "config.json" {
		t.Errorf("Response filename is (%s), expecting (config.json)", testBody.FileName)
	}
	if testBody.Text == "" {
		t.Error("Response file text is blank")
	}
	os.Remove("config.json")
}
