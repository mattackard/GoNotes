package notes

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/mattackard/project-0/pkg/config"
)

func TestCreateFile(t *testing.T) {
	os.Args = []string{"", "File", "Text"}

	//Create a test file
	editFile := false
	CreateFile(config.Mycfg, "testing.txt", &editFile)

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
	if string(text) != "Text" {
		t.Errorf("Text in created file does not match text passed in : %s", string(text))
	}

	Delete("testing.txt")
	os.Remove("./MyNoteFiles")
	Delete("config.json")
}

func ExampleCreateFile() {
	editFile := false
	CreateFile(config.Mycfg, "TestFile.txt", &editFile)
	file, _ := os.Open("TestFile.txt")
	fmt.Println(file != nil)
	//Output: true

	//Delete("TestFile.txt")
	//Delete("config.json")
}

func TestConfig(t *testing.T) {

}

func ExampleConfig() {

}

func TestPrint(t *testing.T) {
	editFile := false
	os.Args = []string{"", "File", "Text"}
	CreateFile(config.Mycfg, "TestFile.txt", &editFile)
	text := Print("TestFile.txt")
	if text != "Text" {
		t.Errorf("File's contents '%s' do not match text given '%s'", text, "Text")
	}

	Delete("TestFile.txt")
	Delete("config.json")
}

func ExamplePrint() {
	editFile := false
	os.Args = []string{"", "File", "Text"}
	CreateFile(config.Mycfg, "TestFile.txt", &editFile)
	text := Print("TestFile.txt")
	fmt.Println(text == "Text")
	//Output: true

	//Delete("TestFile.txt")
	//Delete("config.json")
}

func TestEdit(t *testing.T) {

}

func ExampleEdit() {

}

func TestDelete(t *testing.T) {
	editFile := false
	CreateFile(config.Mycfg, "TestFile.txt", &editFile)
	Delete("TestFile.txt")
	file, _ := os.Open("TestFile.txt")
	if file != nil {
		t.Errorf("The test file could not be deleted")
	}

	Delete("config.json")
}

func ExampleDelete() {
	editFile := false
	CreateFile(config.Mycfg, "TestFile.txt", &editFile)
	Delete("TestFile.txt")
	_, err := os.Open("TestFile.txt")
	fmt.Println(err != nil)
	//Output: true

	//Delete("config.json")
}
