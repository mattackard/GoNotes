
let newNote = document.getElementById("newNote");
let openNote = document.getElementById("openNote");
let deleteNote = document.getElementById("deleteNote");
let saveNote = document.getElementById("saveNote");
let settings = document.getElementById("settings");
let noteEditor = document.getElementById("noteEditor");
let noteTitle = document.getElementById("noteTitle")
let titleInput = document.createElement("INPUT")
titleInput.type = "text";
titleInput.placeholder = "Enter title";

newNote.addEventListener("click", e => {
    e.preventDefault();
    fetch("http://localhost:5555/newNote")
    .then(response => {
        noteTitle.innerText = "";
        noteTitle.insertAdjacentElement("afterbegin", titleInput);
        response.json().then(json => {
            titleInput.value = "";
            noteEditor.value = json.text;
        });
    })
})

openNote.addEventListener("click", e => {
    e.preventDefault();
    noteEditor.value = "Open clicked";
})

deleteNote.addEventListener("click", e => {
    e.preventDefault();
    fetch("http://localhost:5555/deleteNote", {
        method: "POST",
        header: {
            "Content-Type": "text/plain"
        },
        body: JSON.stringify({
            "fileName": noteTitle.innerText,
        })
    })
    .then(response => {
        noteEditor.value = `${noteTitle.innerText} has been deleted.`;
        noteTitle.innerText = "";
    })
})

saveNote.addEventListener("click", e => {
    e.preventDefault();
    let newTitle = "";
    if (titleInput.value == "") {
        newTitle = noteTitle.innerText;
    } else {
        newTitle = titleInput.value;
    }
    titleInput.value = "";
    titleInput.remove();
    noteTitle.innerText = newTitle;
    fetch("http://localhost:5555/saveNote", {
        method: "POST",
        headers: {"Content-Type": "text/plain"},
        body: JSON.stringify({
            "fileName": newTitle,
            "text": noteEditor.value
        })
    })
    .then(response => {
        alert("Save successful");
    })
})

settings.addEventListener("click", e => {
    e.preventDefault();
    fetch("http://localhost:5555/settings")
    .then(response => {
        response.json().then(json => {
            noteTitle.innerText = json.fileName;
            noteEditor.value = json.text;
        });
    })
})