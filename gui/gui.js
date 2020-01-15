
let newNote = document.getElementById("newNote");
let deleteNote = document.getElementById("deleteNote");
let saveNote = document.getElementById("saveNote");
let settings = document.getElementById("settings");
let noteEditor = document.getElementById("noteEditor");

newNote.addEventListener("click", e => {
    e.preventDefault();
    fetch("http://localhost:5555/newNote", {
        mode: "no-cors"
    })
    .then(response => {
       console.log(response);
       noteEditor.innerText = response.body;
    })
})

deleteNote.addEventListener("click", e => {
    e.preventDefault();
    fetch("http://localhost:5555/deleteNote", {
        mode: "no-cors"
    })
    .then(response => {
       noteEditor.innerText = response.body;
    })
})

saveNote.addEventListener("click", e => {
    e.preventDefault();
    fetch("http://localhost:5555/saveNote", {
        mode: "no-cors"
    })
    .then(response => {
       noteEditor.innerText = response.body;
    })
})

settings.addEventListener("click", e => {
    e.preventDefault();
    fetch("http://localhost:5555/settings", {
        mode: "no-cors"
    })
    .then(response => {
       noteEditor.innerText = response.body;
    })
})