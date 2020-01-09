package main

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

func loadConfig() *ini.File {
	//attempt to load the config file
	config, err := ini.Load("config.ini")

	//config error handling
	if err != nil {
		switch err.(type) {

		//if config.ini can't be found, create it
		case *os.PathError:
			println("Found a path error! Creating your .ini file. .")
			config = createNewConfig()
		//if any other error, log it
		default:
			log.Fatal(err)
		}
	}
	return config
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
