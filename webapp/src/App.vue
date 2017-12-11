<template>
  <div id="main">
    <SongList id="songList" @songClicked="loadSong"></SongList>
    <SongDisplay id="songDisplay" :songData="songData"></SongDisplay>
  </div>
</template>

<script>
  import SongList from './components/SongList.vue'
  import SongDisplay from './components/SongDisplay.vue'

  export default {
    name: 'app',
    components: {
      SongList: SongList,
      SongDisplay: SongDisplay,
    },
    data() {
      return {
        songData: ""
      }
    },
    methods: {
      loadSong(songTitle) {
        let xreq = new XMLHttpRequest();
        let vueInst = this;
        xreq.onreadystatechange = function () {
          if (this.readyState == 4 && this.status == 200) {
            let jsonRes = JSON.parse(this.responseText);
            jsonRes.lyrics = parseLyrics(jsonRes.lyrics);
            vueInst.songData = jsonRes;
          }
        }
        xreq.open("GET", "/song/" + songTitle, true);
        xreq.send();
      }
    }
  }

  // Helper functions

  function parseLyrics(lyrics) {
    let lines = lyrics.split("\n")

    let parsed = ""
    for (let line of lines) {
      let parsedLine = ""

      switch(line.slice(0, 1)) {
        case "!":
          parsedLine = parseSection(line) + "\n"
          break;
        case " ":
          parsedLine = line + "\n";
          break;
        case ";":
          parsedLine = line + "\n";
          break;
        case ".":
          parsedLine = " " + line.slice(1) + "\n";
          break;
        default:
          parsedLine = line + "\n";
      }
  
      parsed += parsedLine
    }

    return parsed

    function parseSection(line) {
      let firstLetter = line.slice(1, 2).toLowerCase()
      switch (firstLetter) {
        case "v":
          return "Verse " + line.slice(2);
          break;
        case "c":
          return "Chorus";
          break;
        case "b":
          return "Bridge";
          break;
        case "p":
          return "Pre-Chorus";
          break;
        case "i":
          return "Intro";
          break;
        case "e":
          return "Ending";
          break;
        default:
          return line.slice(1);
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