<template>
  <div id="main">
    <SongList id="songList" @songClicked="loadSong"></SongList>
    <SongDisplay id="songDisplay" :songHtml="songHtml"></SongDisplay>
  </div>
</template>

<script>
  import SongList from './components/SongList.vue'
  import SongDisplay from './components/SongDisplay.vue'
  import decodeSong from './decode.ts'

  export default {
    name: 'app',
    components: {
      SongList: SongList,
      SongDisplay: SongDisplay,
    },
    data() {
      return {
        songHtml: ""
      }
    },
    methods: {
      loadSong(songTitle) {
        let xreq = new XMLHttpRequest();
        let vueInst = this;
        xreq.onreadystatechange = function () {
          if (this.readyState == 4 && this.status == 200) {
            let jsonRes = JSON.parse(this.responseText);
            vueInst.songHtml = decodeSong(jsonRes);
            console.log("")
          }
        }
        xreq.open("GET", "/song/" + songTitle, true);
        xreq.send();
      }
    }
  }

</script>

<style>
  body {
    margin: 0px;
  }

  #main {
    display: flex;
    flex-direction: row;
    height: 100%;
    width: 100%;
  }

  #songList {
    flex-grow: 1;
    width: 30%;
    display: flex;
    flex-flow: column nowrap;
    height: 100vh;
    overflow: auto;
  }

  #songDisplay {
    width: 70%;
    flex-grow: 1;
    height: 100vh;
    overflow: auto;
  }
</style>