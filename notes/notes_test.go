package notes

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mattackard/project-0/config"
)

var cfg config.Config

func TestCreateFile(t *testing.T) {
	cfg = config.LoadConfig()
	os.Args = []string{"", "", "Test text"}

	//Create a test file
	CreateFile(cfg, "testing.txt")

	//Check test file has been created
	_, err := os.Open("testing.txt")
	if err != nil {
		t.Errorf("Created file could not be opened.")
	}
	//Check if text is written to file
	text, err := ioutil.ReadFile("testing.txt")
	if err != nil {
		t.Errorf("Error reading file at testing.txt")
	}
	if string(text) != "Test text" {
		t.Error("Text in created file does not match text passed in")
	}

	//All tests done. Delete test files
	Delete("testing.txt")
	os.Remove("./MyNoteFiles")
}

func TestConfig(t *testing.T) {

}

func TestPrint(t *testing.T) {

}

func TestEdit(t *testing.T) {

}

func TestDelete(t *testing.T) {

}
