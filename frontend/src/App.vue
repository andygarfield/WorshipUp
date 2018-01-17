<template>
  <div id="main">
    <HeaderMenu id="header-menu">
    </HeaderMenu>
    <div id="app-body">
      <Library
        id="library"
        @newSong="newSong">
      </Library>
      <SongDisplay id="song-display"></SongDisplay>
      <ServiceOrder id="service-order"></ServiceOrder>
    </div>
  </div>
</template>

<script>
  import HeaderMenu from './components/HeaderMenu.vue'
  import Library from './components/Library.vue'
  import SongDisplay from './components/SongDisplay.vue'
  import ServiceOrder from './components/ServiceOrder.vue'
  import decodeSong from './decode.ts'

  export default {
    name: 'app',
    components: {
      HeaderMenu: HeaderMenu,
      Library: Library,
      SongDisplay: SongDisplay,
      ServiceOrder: ServiceOrder,
    },
    data() {
      return {
        songHtml: "",
        mode: "read",
        refreshSongs: false,
      }
    },
    methods: {
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

    #library {
      background: #ddd;
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
      padding: 15px;
      flex-basis: 50%;
    }

    #service-order {
      flex-basis: 25%;
      background: #ddd;
    }
  }
</style>