import Vue from 'vue'
import Vuex from 'vuex'
import App from './App.vue'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        mode: "read",
        serviceDate: new Date(),
        songData: "",
        songList: [],
        setLists: [],
        settingsOn: false,
    },
    mutations: {
        switchMode (state, newMode) {
            state.mode = newMode;
        },
        setServiceDate (state, newDate) {
            state.serviceDate = newDate;
        },
        changeSongData (state, songData) {
            state.songData = songData;
        },
        saveSongList (state, songList) {
            state.songList = songList;
        },
        saveSetLists (state, setLists) {
            state.setLists = setLists;
        },
        toggleSettings (state) {
            state.settingsOn = !state.settingsOn;
        },
    },
    actions: {
        loadSong({commit, state}, songTitle) {
            if (state.mode != "read") {
                commit("switchMode", "read");
            }
            
            let xreq = new XMLHttpRequest();
            xreq.onload = function () {
                let jsonRes = JSON.parse(this.responseText);
                commit("changeSongData", jsonRes);
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
        getSetLists ({commit}) {
            let xreq = new XMLHttpRequest();
            xreq.onload = function () {
                commit("saveSetLists", JSON.parse(this.responseText));
            };
            xreq.open("GET", "/setlists/", true);
            xreq.send();
        }
    }
})

new Vue ({
    el: '#app',
    store,
    render: h => h(App)
})
