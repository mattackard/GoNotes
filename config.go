package main

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

var cfg *ini.File

func loadConfig() {
	//attempt to load the config file
	config, err := ini.Load("config.ini")
	cfg = config

	//config error handling
	if err != nil {
		switch err.(type) {

		//if config.ini can't be found, create it
		case *os.PathError:
			println("Found a path error! Creating your .ini file. .")
			cfg = createNewConfig()
		//if any other error, log it
		default:
			log.Fatal(err)
		}
	}
}

func createNewConfig() *ini.File {
	os.Create("config.ini")
	//load config once it as been created
	config, err := ini.Load("config.ini")
	if err != nil {
		log.Fatal(err)
	}
	return config
}
