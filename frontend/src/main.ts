import Vue from 'vue';
import Vuex from 'vuex';
import App from './App.vue';

Vue.use(Vuex)

interface setLists {
    [simpleDate: string]: {
        Date: string,
        Songs: string[],
    }
}

const store = new Vuex.Store({
    state: {
        mode: "read",
        serviceDate: new Date(),
        songData: "",
        songList: [],
        setLists: <setLists> {},
        settingsOn: false,
    },
    getters: {
        stringDate: state => {
            return dateToString(state.serviceDate);
        },
        setSongs: (state, getters) => {
            if (!state.setLists[getters.stringDate]) {
                return [];
            } else {
                return state.setLists[getters.stringDate].Songs;
            }
        }
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
        updateSetList (state, updatedSet) {
            // console.log(updatedSet.stringDate);
            state.setLists[updatedSet.stringDate].Songs = updatedSet.setList;
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

// Helper functions
function dateToString(date: Date) {
    let leftPad = function (str: string, len: number) {
        let sl = str.length;
        if (sl < len) {
            let newString = "";

            let diff = len - sl;
            for (let i = 0; i < diff; i++) {
                newString += "0";
            }
            return newString += str;
        }
        return str;
    }

    let year = leftPad(date.getFullYear().toString(), 2);
    let month = leftPad((date.getMonth()+1).toString(), 2);
    let day = leftPad(date.getDate().toString(), 2);


    return year + month + day;
}