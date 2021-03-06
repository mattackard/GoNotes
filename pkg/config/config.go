// Package config loads and structures the configuration values for GoNotes
package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Mycfg contains the configuration loaded from config.json
// If no cofig.json can be found, the default will be applied
var Mycfg Config

// Default contains a config struct with default values
var Default Config

// Config stores the path and options for GoNotes
type Config struct {
	Paths   Path   `json:"paths"`
	Options Option `json:"options"`
}

// Path stores filepath information
type Path struct {
	Notes string `json:"notes"`
}

// Option stores user defined options for toggling datestamps, eiditor on initialization, file extension and server port
type Option struct {
	DateStamp     bool   `json:"dateStamp"`
	InitEditor    bool   `json:"initEditor"`
	FileExtension string `json:"fileExtension"`
	Port          string `json:"port"`
}

// FullPath holds the path including directory to use when referencing note files
// var FullPath string

// initializes the global variables having to do with configuration of the program
func init() {

	Default = Config{
		Paths: Path{
			Notes: "./notes/",
		},
		Options: Option{
			DateStamp:     true,
			InitEditor:    false,
			FileExtension: ".txt",
			Port:          ":6060",
		},
	}
	Mycfg = LoadConfig()
}

// LoadConfig loads the ./config.json and parses it into the Config struct
func LoadConfig() (cfg Config) {
	// attempt to load the config file
	jsonFile, err := os.Open("./config.json")

	// config error handling
	if err != nil {
		switch err.(type) {

		// if config.ini can't be found, create it
		case *os.PathError:
			println("Config file could not be found. Creating your config.json file. . .")
			jsonFile = createNewConfig()
		// if any other error, log it
		default:
			panic(err)
		}
	}
	bytes, err := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &cfg)
	return cfg
}

// CreateNewConfig generates a new config.json file at ./config.json and loads in the default values
func createNewConfig() *os.File {
	os.Create("./config.json")
	file, err := json.MarshalIndent(Default, "", "    ")
	err = ioutil.WriteFile("config.json", file, 0666)
	if err != nil {
		panic(err)
	}
	// load config once it as been created
	config, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	return config
}
