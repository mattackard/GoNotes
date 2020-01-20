package main

import "github.com/zserge/webview"

//open a webview connection using the client.html file
//css and js is loaded in through the js
func main() {
	webview.Open("GoNotes", "https://raw.githack.com/mattackard/project-0/master/cmd/GoNotesClient/client.html", 600, 700, true)
}
