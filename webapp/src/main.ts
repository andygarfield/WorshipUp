import Vue from 'vue'
import App from './App.vue'


new Vue({
    el: '#app',
    render: h => h(App)
})

// interface SongJSON {
//     title: string
//     lyrics: string
//     presentation?: string
//     author?: string
//     ccli?: string
// }

// var song: SongJSON;
// var songList: string[];

// var xreq = new XMLHttpRequest();
// xreq.onreadystatechange = function() {
//     if (this.readyState == 4 && this.status == 200) {
//         songList = JSON.parse(this.responseText);
//         renderList(songList);
//     }
// }
// xreq.open("GET", "/songlist/", true);
// xreq.send();

// function renderList(songList: string[]) {
//     let target = document.getElementById("app");
//     if (target) {
//         let newList = document.createElement("ul")
//         for (let title of songList) {
//             let newItem = document.createElement("li");
//             newItem.innerText = title;
//             newList.appendChild(newItem);
//         }
//         target.appendChild(newList);
//     }
// }

// function renderSong(song: SongJSON) {
//     let target = document.getElementById("app");
//     if (target) {
//         let songTitle = document.createElement("h3");
//         songTitle.innerText = song.title;
//         target.appendChild(songTitle);

//         let paragraph = document.createElement("p");
//         paragraph.innerText = song.lyrics;
//         target.appendChild(paragraph);
//     }
// }