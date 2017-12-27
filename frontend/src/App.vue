<template>
  <div id="main">
    <HeaderMenu id="header-menu">
    </HeaderMenu>
    <div id="app-body">
      <Library
        id="library"
        @songClicked="loadSong"
        @newSong="newSong">
      </Library>
      <SongDisplay id="song-display" :songHtml="songHtml" :mode="mode"></SongDisplay>
    </div>
  </div>
</template>

<script>
  import HeaderMenu from './components/HeaderMenu.vue'
  import Library from './components/Library.vue'
  import SongDisplay from './components/SongDisplay.vue'
  import decodeSong from './decode.ts'

  export default {
    name: 'app',
    components: {
      HeaderMenu: HeaderMenu,
      Library: Library,
      SongDisplay: SongDisplay,
    },
    data() {
      return {
        songHtml: "",
        mode: "read",
      }
    },
    methods: {
      loadSong(songTitle) {
        this.mode = "read";

        let xreq = new XMLHttpRequest();
        let vueInst = this;
        xreq.onreadystatechange = function () {
          if (this.readyState == 4 && this.status == 200) {
            let jsonRes = JSON.parse(this.responseText);
            vueInst.songHtml = decodeSong(jsonRes);
          }
        }
        xreq.open("GET", "/song/" + songTitle, true);
        xreq.send();
      },
      newSong() {
        this.mode = "edit";
      },
    }
  }
</script>

<style>
  @media (min-width: 0px) {
    body {
      font-family: 'Fira Sans', sans-serif;
      margin: 0px;
    }

    #header-menu {
      height: 15vh;
    }
  }

  @media (min-width: 690px) {
    #app-body {
      display: flex;
      flex-direction: row;
      height: 85vh;
    }

    #app-body>* {
      overflow: auto;
    }

    #library {
      background: #ddd;
      flex-basis: 25%;
    }

    #library>* {
      cursor: pointer;
    }

    #song-display {
      /* margin-left: 6px; */
      flex-basis: 75%;
    }

    #service-order {}
  }
</style>