//set up variables for all elements that will be accessed
let newNote = document.getElementById("newNote");
let openNote = document.getElementById("openNote");
let deleteNote = document.getElementById("deleteNote");
let saveNote = document.getElementById("saveNote");
let settings = document.getElementById("settings");
let noteEditor = document.getElementById("noteEditor");
let noteTitle = document.getElementById("noteTitle");
let files = document.getElementById("files");
let fileBrowser = document.getElementById("fileBrowser");

//create input element for new file titles
let titleInput = document.createElement("INPUT");
titleInput.type = "text";
titleInput.placeholder = "Enter title";

//initiailize working directory to project directory
let workingDir = "./";

//removes the file path in front of filename
let parseTitle = path => {
  let pathArr = path.split("/");
  let title = pathArr[pathArr.length - 1];
  title = title.split(".");
  if (title[1] == "txt") {
    return title[0];
  } else {
    return pathArr[pathArr.length - 1];
  }
};

newNote.addEventListener("click", e => {
  e.preventDefault();
  //clear the file broswer list and hide it behind other content again
  files.innerHTML = "";
  fileBrowser.style.zIndex = -1;
  fetch("http://localhost:5555/newNote").then(response => {
    //reset note title input value and put it into the html
    noteTitle.innerText = "";
    noteTitle.insertAdjacentElement("afterbegin", titleInput);

    //parse JSON response into title and editor fields
    response.json().then(json => {
      titleInput.value = "";
      noteEditor.value = json.text;
    });
  });
});

openNote.addEventListener("click", e => {
  e.preventDefault();

  //clears note editor text
  noteEditor.value = "";
  titleInput.remove();
  noteTitle.innerText = "Open File";

  fetchDir(workingDir);
});

let fetchDir = dir => {
  fetch("http://localhost:5555/dir", {
    method: "POST",
    header: {
      "Content-Type": "text/plain"
    },
    body: JSON.stringify({
      root: dir,
      files: ""
    })
  }).then(response => {
    //parses the response to json and then creates a list element for each file
    response.json().then(json => {
      //adds the path to move up the file tree
      files.innerHTML += `<li class="file" title="../">
                                    <svg height="60" version="1.1" width="100" viewBox="3 2 20 20" ><g transform="translate(0 -1028.4)" title="../"><path d="m12 1034.4c0 1.1-0.895 2-2 2h-5-3c-1.1046 0-2 0.9-2 2v8 3c0 1.1 0.89543 2 2 2h20c1.105 0 2-0.9 2-2v-3-10c0-1.1-0.895-2-2-2h-10z" fill="#2980b9""/><path d="m2 2c-1.1046 0-2 0.8954-2 2v5h10v1h14v-5c0-1.1046-0.895-2-2-2h-10.281c-0.346-0.5969-0.979-1-1.719-1h-8z" fill="#2980b9" transform="translate(0 1028.4)"/><path d="m12 1033.4c0 1.1-0.895 2-2 2h-5-3c-1.1046 0-2 0.9-2 2v8 3c0 1.1 0.89543 2 2 2h20c1.105 0 2-0.9 2-2v-3-10c0-1.1-0.895-2-2-2h-10z" fill="#3498db"/></g></svg>
                                    <p title="../">../</p>
                                </li>`;
      json.files.forEach(file => {
        //create html string to append to empty list
        let elem = "";

        //if the file name has a ".", assume the file is a folder and give it a folder svg, otherwise give it the document svg
        if (file.includes(".")) {
          elem = `<li class="file" title="${file}">
                                <svg height="60" version="1.1" width="100" viewBox="3 2 20 20" ><g transform="translate(0 -1028.4)"><path d="m5 1030.4c-1.1046 0-2 0.9-2 2v8 4 6c0 1.1 0.8954 2 2 2h14c1.105 0 2-0.9 2-2v-6-4-4l-6-6h-10z" fill="#95a5a6"/><path d="m5 1029.4c-1.1046 0-2 0.9-2 2v8 4 6c0 1.1 0.8954 2 2 2h14c1.105 0 2-0.9 2-2v-6-4-4l-6-6h-10z" fill="#bdc3c7"/><path d="m21 1035.4-6-6v4c0 1.1 0.895 2 2 2h4z" fill="#95a5a6"/><path d="m6 8v1h12v-1h-12zm0 3v1h12v-1h-12zm0 3v1h12v-1h-12zm0 3v1h12v-1h-12z" fill="#95a5a6" transform="translate(0 1028.4)"/></g></svg>               
                                <p title="${file}">${file}</p>
                            </li>`;
        } else {
          elem = `<li class="file" title="${file}/">
                                <svg height="60" version="1.1" width="100" viewBox="3 2 20 20" ><g transform="translate(0 -1028.4)"><path d="m12 1034.4c0 1.1-0.895 2-2 2h-5-3c-1.1046 0-2 0.9-2 2v8 3c0 1.1 0.89543 2 2 2h20c1.105 0 2-0.9 2-2v-3-10c0-1.1-0.895-2-2-2h-10z" fill="#2980b9"/><path d="m2 2c-1.1046 0-2 0.8954-2 2v5h10v1h14v-5c0-1.1046-0.895-2-2-2h-10.281c-0.346-0.5969-0.979-1-1.719-1h-8z" fill="#2980b9" transform="translate(0 1028.4)"/><path d="m12 1033.4c0 1.1-0.895 2-2 2h-5-3c-1.1046 0-2 0.9-2 2v8 3c0 1.1 0.89543 2 2 2h20c1.105 0 2-0.9 2-2v-3-10c0-1.1-0.895-2-2-2h-10z" fill="#3498db"/></g></svg>
                                <p title="${file}/">${file}/</p>
                            </li>`;
        }
        files.innerHTML += elem;
      });
      //after files have been added to list, bring list into view
      fileBrowser.style.zIndex = 2;
    });
  });
};

files.addEventListener("click", e => {
  //if target has a / assume it is a directory and get that directories content
  //otherwise get the file clicked
  if (e.target.title.includes("/")) {
    //clear existing files from list
    files.innerHTML = "";

    //add new directory onto working directory path
    workingDir += e.target.title;
    fetchDir(workingDir);
  } else {
    fetch("http://localhost:5555/getFile", {
      method: "POST",
      header: {
        "Content-Type": "text/plain"
      },
      body: JSON.stringify({
        path: workingDir,
        fileName: e.target.title,
        text: ""
      })
    }).then(response => {
      //clear the file broswer list and hide it behind other content again
      files.innerHTML = "";
      fileBrowser.style.zIndex = -1;

      //parse file contents into title and editor
      response.json().then(json => {
        noteTitle.innerText = parseTitle(json.fileName);
        noteEditor.value = json.text;
      });
    });
  }
});

deleteNote.addEventListener("click", e => {
  //clear the file broswer list and hide it behind other content again
  files.innerHTML = "";
  fileBrowser.style.zIndex = -1;

  e.preventDefault();
  //send filename to delete file
  fetch("http://localhost:5555/deleteNote", {
    method: "POST",
    header: {
      "Content-Type": "text/plain"
    },
    body: JSON.stringify({
      path: workingDir,
      fileName: noteTitle.innerText,
      text: ""
    })
  }).then(response => {
    //put confirmation message into editor area and clear title
    noteEditor.value = `${noteTitle.innerText} has been deleted.`;
    noteTitle.innerText = "";
  });
});

saveNote.addEventListener("click", e => {
  e.preventDefault();

  //clear the file broswer list and hide it behind other content again
  files.innerHTML = "";
  fileBrowser.style.zIndex = -1;

  //if the title field is an input, get the value
  //otherwise get the value of the title field
  let newTitle = "";
  if (titleInput.value == "") {
    newTitle = noteTitle.innerText;
  } else {
    newTitle = titleInput.value;
  }

  //reset the input value for next new file and remove the element
  titleInput.value = "";
  titleInput.remove();
  noteTitle.innerText = newTitle;
  fetch("http://localhost:5555/saveNote", {
    method: "POST",
    headers: { "Content-Type": "text/plain" },
    body: JSON.stringify({
      path: workingDir,
      fileName: newTitle,
      text: noteEditor.value
    })
  });
});

//get the config file put it's contents into the editor
settings.addEventListener("click", e => {
  //clear the file broswer list and hide it behind other content again
  files.innerHTML = "";
  fileBrowser.style.zIndex = -1;

  e.preventDefault();
  fetch("http://localhost:5555/settings").then(response => {
    response.json().then(json => {
      noteTitle.innerText = parseTitle(json.fileName);
      noteEditor.value = json.text;
    });
  });
});
