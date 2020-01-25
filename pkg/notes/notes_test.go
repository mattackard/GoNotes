package notes

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/mattackard/project-0/pkg/config"
)

func TestCreateFile(t *testing.T) {
	os.Args = []string{"File", "Hello", "World"}

	// Create a test file
	CreateFile(config.Mycfg.Paths.Notes, "testing.txt", false, false)

	// Check test file has been created
	test, err := os.Open("testing.txt")
	defer test.Close()
	if err != nil {
		t.Errorf("Created file could not be opened.")
	}

	// Check if text is written to file
	text, err := ioutil.ReadFile("testing.txt")
	if err != nil {
		t.Errorf("Error reading file at testing.txt")
	}
	if string(text) != "Hello World " {
		t.Errorf("Text in created file does not match text passed in : %s", string(text))
	}

	Delete("testing.txt")
	os.Remove("./MyNoteFiles")
	Delete("config.json")
}

func ExampleCreateFile() {
	CreateFile(config.Mycfg.Paths.Notes, "TestFile.txt", false, false)
	file, _ := os.Open("TestFile.txt")
	defer file.Close()
	fmt.Println(file != nil)
	// Output: true

	// Delete("TestFile.txt")
	// Delete("config.json")
}

func TestPrint(t *testing.T) {
	os.Args = []string{"File", "Text"}
	CreateFile(config.Mycfg.Paths.Notes, "TestFile.txt", false, false)
	text := Print("TestFile.txt")
	if text != "Text " {
		t.Errorf("File's contents '%s' do not match text given '%s'", text, "Text ")
	}

	Delete("TestFile.txt")
	Delete("config.json")
}

func ExamplePrint() {
	os.Args = []string{"File", "Text"}
	CreateFile(config.Mycfg.Paths.Notes, "TestFile.txt", false, false)
	text := Print("TestFile.txt")
	fmt.Println(text == "Text")
	// Output: true

	// Delete("TestFile.txt")
	// Delete("config.json")
}

func TestDelete(t *testing.T) {
	CreateFile(config.Mycfg.Paths.Notes, "TestFile.txt", false, false)
	Delete("TestFile.txt")
	file, _ := os.Open("TestFile.txt")
	defer file.Close()
	if file != nil {
		t.Errorf("The test file could not be deleted")
	}

	Delete("config.json")
}

func ExampleDelete() {
	CreateFile(config.Mycfg.Paths.Notes, "TestFile.txt", false, false)
	Delete("TestFile.txt")
	test, err := os.Open("TestFile.txt")
	defer test.Close()
	fmt.Println(err != nil)
	// Output: true

	// Delete("config.json")
}
