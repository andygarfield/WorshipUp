import Vue from 'vue';
import Vuex from 'vuex';
import App from './App.vue';
import Promise from 'promise-polyfill';

export { SongJSON, SetList, IndexedSong };
Vue.use(Vuex);

interface SongJSON {
    id: string
    title: string
    body?: string
    presentation?: string
    author?: string
    ccli?: string
}

interface IndexedSong {
    i: number
    song: SongJSON
}

interface SetList {
    ID: string
    Date: Date
    Songs: string[]
}

interface state {
    mode: string
    songList: IndexedSong[]
    displayedSong: IndexedSong
    // currentIndex: number
    displayedSetList: SetList
    serviceDate: Date
    setLists: SetList[]
    settingsOn: boolean
}

let blankSong = <IndexedSong> {
    i: -1,
    song: {
        id: "-1",
        title: '',
        body: '',
    },
};

const store = new Vuex.Store({
    state: <state> {
        mode: "read",

        displayedSong: blankSong,
        songList: [] as IndexedSong[],

        displayedSetList: {
            ID: "-1",
            Date: new Date(),
        },
        setLists: [] as SetList[],
        settingsOn: false,
    },
    getters: {
        stringDate: state => {
            return dateToString(state.displayedSetList.Date);
        },
        // currentIndex: state => {
        //     let currentIndex = -1;
        //     state.songList.forEach((song) => {
        //         if (state.displayedSong.id === song.id) {
        //             currentIndex++;
        //         }
        //         return currentIndex;
        //     })
        // },
        // setListSongs: (state, getters) => {
        //     if (!state.setLists[getters.stringDate]) {
        //         return [];
        //      }else {
        //         return state.setLists[getters.stringDate].Songs;
        //     }
        // }
    },
    mutations: {
        switchMode (state, newMode: string) {
            state.mode = newMode;
        },
        saveSongList (state, songList: SongJSON[]) {
            songList.sort((a, b) => {
                let titleA = a.title.toLowerCase();
                let titleB = b.title.toLowerCase();
                if (titleA < titleB) {
                    return -1;
                }
                if (titleA > titleB) {
                    return 1;
                }
                
                return 0;
            });
            state.songList = songList.map((item, i) => {
                return <IndexedSong> {
                    i: i,
                    song: item,
                }
            })
        },
        changeDisplayedSong (state, song: IndexedSong) {
            state.displayedSong = song;
        },
        makeSongBlank (state) {
            state.displayedSong = blankSong;
            // state.currentIndex = -1;
        },
        // changeSongIndex (state, index: number) {
        //     state.currentIndex = index;
        // },
        saveSetLists (state, setLists) {
            state.setLists = setLists;
            // state.songList.sort();
        },
        // sortAlpha (state) {
        //     // state.songList.sort();
        //     state.songList.sort((a, b) => {
        //         let titleA = a.title.toLowerCase();
        //         let titleB = b.title.toLowerCase();
        //         if (titleA < titleB) {
        //             return -1;
        //         }
        //         if (titleA > titleB) {
        //             return 1;
        //         }
                
        //         return 0;
        //         });
        // },
        // updateSetList (state, updatedSet) {
        //     state.setLists[updatedSet.stringDate].Songs = updatedSet.setList;
        // },
        toggleSettings (state) {
            state.settingsOn = !state.settingsOn;
        },
    },
    actions: {
        getSong({commit, state}, iSong: IndexedSong) {
            if (state.mode != "read") {
                commit("switchMode", "read");
            }
            let query = `{song (id: ${iSong.song.id}) { id title body }}`;
            sendData(query, function (this: XMLHttpRequest) {
                let song = JSON.parse(this.responseText).data.song as SongJSON;
                commit("changeDisplayedSong", {id: iSong.i, song: song});
            });
        },
        nextSong({commit, state, dispatch}) {
            let newID: string;
            let newIndex: number;
            let noSongs = state.songList.length == 0;
            let onLastSong = state.displayedSong.i == state.songList.length && !noSongs;

            if (onLastSong) {
                newIndex = state.displayedSong.i-1;
                // commit("changeSongIndex", newIndex);
            } else if (noSongs) {
                state.displayedSong = blankSong;
                state.displayedSong.i = -1;
                return;
            } else {
                newIndex = state.displayedSong.i;
            }
            newID = state.songList[newIndex].song.id;
            dispatch("getSong", newID);
        },
        addSong({dispatch, commit}, song: SongJSON) {
            let query = `
            mutation {
                addsong(song: {
                    title: "${song.title}"
                    body: "${song.body}"
                })
                {
                    id
                    title
                    body
                }
            }
            `;

            sendData(query, function (this: XMLHttpRequest) {
                let song = JSON.parse(this.responseText).data.addsong;
                dispatch("getSongList")
                commit("changeDisplayedSong", song);
                commit("switchMode", "read");
            })
        },
        updateSong({dispatch, commit}, song: SongJSON) {
            let query = `
            mutation {
                updatesong(
                id: ${song.id}
                song: {
                    title: "${song.title}"
                    body: "${song.body}"
                })
                {
                    id
                    title
                    body
                }
            }
            `;

            sendData(query, function (this: XMLHttpRequest) {
                let song = JSON.parse(this.responseText).data.updatesong;
                // dispatch("getSongList")
                commit("changeDisplayedSong", song);
                commit("switchMode", "read");
            })
        },
        deleteSong({state, dispatch}, id) {
            let query = `
            mutation {
                deletesong(id: ${id})
            }
            `;
            sendData(query, function(this: XMLHttpRequest) {
                dispatch("getSongList")
                .then(() => {
                    dispatch("nextSong");
                })
            });
	    },
        getSongList ({commit}) {
            return new Promise((resolve, reject) => {
                let query = '{allsongs { id title }}'
                sendData(query, function (this: XMLHttpRequest) {
                    commit("saveSongList", JSON.parse(this.responseText).data.allsongs);
                    resolve();
                });
            })
        },

        getSet() {},
        addSet() {},
        updateSet() {},
        deleteSet() {},
        getAllSets ({commit}) {
            let query = '{allsets { id date }}'
            sendData(query, function (this: XMLHttpRequest) {
                commit("saveSetLists", JSON.parse(this.responseText).data.allsets);
            });
        },
    },
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

// For handling api calls
function sendData(query: string, callback: () => void) {
    let xreq = new XMLHttpRequest();
    xreq.onload = callback;
    xreq.open("POST", "/gql", true);
    query = JSON.stringify({query: query});
    xreq.send(query);
}