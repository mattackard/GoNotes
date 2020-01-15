
let newNote = document.getElementById("newNote");
let deleteNote = document.getElementById("deleteNote");
let saveNote = document.getElementById("saveNote");
let settings = document.getElementById("settings");
let noteEditor = document.getElementById("noteEditor");

newNote.addEventListener("click", e => {
    e.preventDefault();
    fetch("http://localhost:5555/newNote")
        .then(response => {
        response.json().then(json => {
            noteEditor.innerText = json.text;
        });
    })
})

deleteNote.addEventListener("click", e => {
    e.preventDefault();
    fetch("http://localhost:5555/deleteNote")
    .then(response => {
        response.json().then(json => {
            noteEditor.innerText = json.text;
            alert(`${json.fileName} has been deleted.`)
        })
    })
})

saveNote.addEventListener("click", e => {
    e.preventDefault();
    fetch("http://localhost:5555/saveNote")
    .then(response => {
        response.json().then(json => {
            noteEditor.innerText = json.text;
        });
    })
})

settings.addEventListener("click", e => {
    e.preventDefault();
    fetch("http://localhost:5555/settings")
    .then(response => {
        response.json().then(json => {
            noteEditor.innerText = json.text;
        });
    })
})