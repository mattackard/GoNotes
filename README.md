# GoNotes

GoNotes is a lightweight text editor written in Go. GoNotes was originally written as a CLI application, but has been expanded to include a HTTP client and Go server. The original CLI program is still fully functional and has all the core features of the web client.

## Getting started

To get started working with the project clone the repository

```
git clone https://github.com/mattackard/project-0.git
```

or you can use go get

```
go get github.com/mattackard/project-0
```

## Porject Dependencies

You can install all go dependencies using go get in the project directory

```
go get -u -v -f all
```

The GoNotes Client uses webview which has some dependencies of its on that can be found in the webview [readme](https://github.com/zserge/webview/blob/master/README.md).

## Using GoNotes

GoNotes comes packaged with a local CLI app, and a client and server for the GUI. The go files to run each program can be found within the cmd directory in the project root.

To run the CLI, open a terminal in the cmd/GoNotes directory and run cli.go

```
~/projectDir $ go run cmd/GoNotes/cli.go
```

To run the GUi client you will need to start the server and client applications

```
~/projectDir $ go run cmd/GoNotesd/server.go
```

and in a new terminal instance

```
~/projectDir $ go run cmd/GoNotesClient/client.go
```

## Running the tests

To run the test files for the whole project cd into the project directory and run

```
~/projectDir $ go test ./...
```

## API

The GoNotesd server hosts an API on your localhost:5555.

#### GET /newNote

Returns a JSON object containing the text header added to all new files. The filename field is left blank and set when the request is made to save the file.

```
{
    "fileName": "",
    "text":     "Response text here",
}
```

#### GET /settings

Returns a JSON object with the contents of the user's config.json file.

```
{
    "paths": {
        "notes": "./"
    },
    "options": {
        "dateStamp": true,
        "initEditor": false,
        "fileExtension": ".txt",
        "port": ":5555"
    }
}
```

#### POST /getFile

Takes a JSON object in the following format

```
{
    "path" : "./MyDirectory/",
    "fileName": "MyFile.txt",
}
```

and returns a reponse containing the file path, file name, and text content

```
{
    "path":     "./MyDirectory/",
    "fileName": "MyFile.txt",
    "text":     "Text content will be here",
}
```

#### POST /deleteNote

Takes a JSON object with the following format

```
{
    "path":     "./MyDirectory/",
    "fileName": "MyFile.txt",
}
```

and returns an "OK" message on successful deletion

#### POST /saveNote

Takes a JSON object with the following format

```
{
    "path":     "./MyDirectory/",
    "fileName": "MyFile.txt",
    "text":     "Text content will be here",
}
```

and returns an "OK" message on successful file save

#### POST /dir

Takes a JSON object in the following format

```
{
    "root": dir,
    "files": "",
}
```

and returns the same JSON object back with the files

```

{
    "root": dir,
    "files": ["MyFile.txt", "MyOtherFile.ini"],
}

```

## Built With

- [webview](https://github.com/zserge/webview) - The file server used for the GUI
