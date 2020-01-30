
build:
	go build ./cmd/GoNotes/.
	go build ./cmd/GoNotesd/.
	go build ./cmd/GoNotesClient/.
	mv GoNotes ./bin/
	mv GoNotesd ./bin/
	mv GoNotesClient ./bin/

server: build
	./bin/GoNotesd

client: build
	./bin/GoNotesClient
