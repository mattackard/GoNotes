
let newNote = document.getElementById("newNote");
let openNote = document.getElementById("openNote");
let deleteNote = document.getElementById("deleteNote");
let saveNote = document.getElementById("saveNote");
let settings = document.getElementById("settings");
let noteEditor = document.getElementById("noteEditor");
let noteTitle = document.getElementById("noteTitle");
let titleInput = document.createElement("INPUT");
let fileBrowser = document.getElementById("fileBrowser");
let files = document.getElementById("files");
fileBrowser.style.zIndex = -1;
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
});

openNote.addEventListener("click", e => {
    e.preventDefault();
    noteEditor.value = "";
    fetch("http://localhost:5555/dir")
    .then(response => {
        response.json().then(json => {
            json.files.forEach(file => {
                let elem = "";
                if (file.includes(".")) {
                    elem = `<li class="file" title="${file}">
                                <svg height="60" version="1.1" width="100" viewBox="3 2 20 20" ><g transform="translate(0 -1028.4)" title="${file}"><path d="m5 1030.4c-1.1046 0-2 0.9-2 2v8 4 6c0 1.1 0.8954 2 2 2h14c1.105 0 2-0.9 2-2v-6-4-4l-6-6h-10z" fill="#95a5a6" title="${file}"/><path d="m5 1029.4c-1.1046 0-2 0.9-2 2v8 4 6c0 1.1 0.8954 2 2 2h14c1.105 0 2-0.9 2-2v-6-4-4l-6-6h-10z" fill="#bdc3c7" title="${file}"/><path d="m21 1035.4-6-6v4c0 1.1 0.895 2 2 2h4z" fill="#95a5a6" title="${file}"/><path d="m6 8v1h12v-1h-12zm0 3v1h12v-1h-12zm0 3v1h12v-1h-12zm0 3v1h12v-1h-12z" fill="#95a5a6" transform="translate(0 1028.4)" title="${file}"/></g></svg>               
                                <p title="${file}">${file}</p>
                            </li>`;
                } else {
                    elem = `<li class="file" title="${file}/">
                                <svg height="60" version="1.1" width="100" viewBox="3 2 20 20" ><g transform="translate(0 -1028.4)" title="${file}"><path d="m12 1034.4c0 1.1-0.895 2-2 2h-5-3c-1.1046 0-2 0.9-2 2v8 3c0 1.1 0.89543 2 2 2h20c1.105 0 2-0.9 2-2v-3-10c0-1.1-0.895-2-2-2h-10z" fill="#2980b9" title="${file}"/><path d="m2 2c-1.1046 0-2 0.8954-2 2v5h10v1h14v-5c0-1.1046-0.895-2-2-2h-10.281c-0.346-0.5969-0.979-1-1.719-1h-8z" fill="#2980b9" transform="translate(0 1028.4)" title="${file}"/><path d="m12 1033.4c0 1.1-0.895 2-2 2h-5-3c-1.1046 0-2 0.9-2 2v8 3c0 1.1 0.89543 2 2 2h20c1.105 0 2-0.9 2-2v-3-10c0-1.1-0.895-2-2-2h-10z" fill="#3498db" title="${file}"/></g></svg>
                                <p title="${file}/">${file}/</p>
                            </li>`;
                }
                files.innerHTML += elem;
            });
            fileBrowser.style.zIndex = 2;
        });
    });
});

files.addEventListener("click", e => {
    if (e.target.title.includes("/")) {
        alert("changing directories is not currently supported");
    } else {
        fetch("http://localhost:5555/getFile", {
            method: "POST",
            header: {
                "Content-Type": "text/plain"
            },
            body: JSON.stringify({
                "fileName": e.target.title,
                "text": "",
            })
        })
        .then(response => {
            files.innerHTML = "";
            fileBrowser.style.zIndex = -1;
            response.json().then(json => {
                noteTitle.innerText = json.fileName;
                noteEditor.value = json.text;
            });
        });
    }
});

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
    });
});

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
    });
});

settings.addEventListener("click", e => {
    e.preventDefault();
    fetch("http://localhost:5555/settings")
    .then(response => {
        response.json().then(json => {
            noteTitle.innerText = json.fileName;
            noteEditor.value = json.text;
        });
    });
});