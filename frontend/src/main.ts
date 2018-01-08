import Vue from 'vue'
import Vuex from 'vuex'
import App from './App.vue'
import { decodeSong } from './decode'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        mode: "read",
        songData: "",
        songList: [],
    },
    mutations: {
        switchMode (state, newMode) {
            state.mode = newMode;
        },
        loadSongData (state, songData) {
            state.songData = songData;
        },
        saveSongList (state, songList) {
            state.songList = songList;
        }
    },
    actions: {
        loadSong({commit, state}, songTitle) {
            if (state.mode != "read") {
                commit("switchMode", "read");
            }
            
            let xreq = new XMLHttpRequest();
            xreq.onload = function () {
                let jsonRes = JSON.parse(this.responseText);
                commit("loadSongData", decodeSong(jsonRes));
            }
            xreq.open("GET", "/song/" + songTitle, true);
            xreq.send();
        },
        getSongList ({commit}) {
            let xreq = new XMLHttpRequest();
            xreq.onload = function () {
                commit("saveSongList", JSON.parse(this.responseText));
            };
            xreq.open("GET", "/songlist/", true);
            xreq.send();
        },
    }
})

new Vue ({
    el: '#app',
    store,
    render: h => h(App)
})
