package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

//Config stores the path and options for GoNotes
type Config struct {
	Paths   Path   `json:"paths"`
	Options Option `json:"options"`
}

//Path stores filepath information
type Path struct {
	Notes string `json:"notes"`
}

//Option stores user defined options for program functionality
type Option struct {
	DateStamp     bool   `json:"dateStamp"`
	FileExtension string `json:"fileExtension"`
}

//LoadConfig loads the ./config.json and parses it into the Config struct
func LoadConfig() (cfg Config) {
	//make sure the os is referencing the project root
	goPath := os.Getenv("GOPATH")
	projectDir := path.Join(goPath, "src/github.com/mattackard/project-0/")

	//attempt to load the config file
	jsonFile, err := os.Open(projectDir + "/config.json")
	defer jsonFile.Close()

	//config error handling
	if err != nil {
		switch err.(type) {

		//if config.ini can't be found, create it
		case *os.PathError:
			println("Found a path error! Creating your config.json file. . .")
			jsonFile = CreateNewConfig()
		//if any other error, log it
		default:
			log.Fatal(err)
		}
	}
	bytes, err := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &cfg)
	return cfg
}

//CreateNewConfig generates a new config.json file at ./config.json
func CreateNewConfig() *os.File {
	os.Create("config.json")
	newConfig := Config{
		Paths: Path{
			Notes: "./",
		},
		Options: Option{
			DateStamp:     true,
			FileExtension: ".txt",
		},
	}
	file, err := json.MarshalIndent(newConfig, "", "    ")
	err = ioutil.WriteFile("config.json", file, 0777)
	if err != nil {
		log.Fatal(err)
	}
	//load config once it as been created
	config, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	return config
}
