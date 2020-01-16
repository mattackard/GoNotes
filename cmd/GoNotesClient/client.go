package main

import "github.com/zserge/webview"

func main() {
	webview.Open("GoNotes", "file:///home/ubuntu/go/src/github.com/mattackard/project-0/cmd/GoNotesClient/client.html", 600, 700, true)
}
