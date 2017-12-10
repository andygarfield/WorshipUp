'use strict'

interface SongJSON {
    title: string
    lyrics: string
    presentation?: string
    author?: string
    ccli?: string
}

var song: SongJSON;

var xreq = new XMLHttpRequest();
xreq.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
        song = JSON.parse(this.responseText);
        renderSong(song);
    }
}
xreq.open("GET", "/songlist/", true);
xreq.send();

function renderSong(song: SongJSON) {
    let target = document.getElementById("target");
    if (target) {
        let songTitle = document.createElement("h3");
        songTitle.innerText = song.title;
        target.appendChild(songTitle);
    
        let paragraph = document.createElement("p");
        paragraph.innerText = song.lyrics;
        target.appendChild(paragraph);
    }


    // function renderLyrics() {
        
    // }
}