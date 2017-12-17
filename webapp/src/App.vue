<template>
  <div id="main">
    <HeaderMenu id="header-menu">
    </HeaderMenu>
    <div id="app-body">
      <Library id="library" @songClicked="loadSong"></Library>
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
      }
    }
  }
</script>

<style>
  @media (min-width: 0px) {
    body {
      font-family: "Trebuchet MS", Helvetica, sans-serif;
      margin: 0px;
    }

 #header-menu {
      height: 20vh;
    }
  }

  @media (min-width: 690px) {
    #app-body {
      display: flex;
      flex-direction: row;
      height: 80vh;
      width: 100vh;
    }

    #app-body > * {
      overflow: auto;
    }

    #library {
      background: #ddd;
    }
    
    #library > * {
      cursor: pointer;
    }

    #song-display {
      flex-grow: 1;
    }

    #service-order {}
  }
</style>