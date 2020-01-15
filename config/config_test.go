package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateNewConfig(t *testing.T) {
	var generatedConfig Config
	newConfig := createNewConfig()
	if newConfig == nil {
		t.Error("Error generating new config file")
	}
	bytes, _ := ioutil.ReadAll(newConfig)
	json.Unmarshal(bytes, &generatedConfig)
	if generatedConfig != Default {
		t.Error("Generated config does not match default config")
	}

	os.Remove("config.json")
}

func TestLoadConfig(t *testing.T) {
	testConfig := LoadConfig()
	if testConfig != Default {
		fmt.Print(testConfig, Default)
		t.Error("Loaded config file's contents do not match the default config")
	}

	os.Remove("config.json")
}

func ExampleLoadConfig() {
	myCfg := LoadConfig()
	println(myCfg.Paths.Notes)               // "./"
	fmt.Println(myCfg.Options.FileExtension) // ".txt"
	//Output: .txt

	//os.Remove("config.json")
}
