<template>
  <div id="main">
    <Settings v-if="settingsOn"></Settings>
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
  import Settings from  './components/Settings.vue'
  import decodeSong from './decode.ts'
  import { mapState } from 'vuex';

  export default {
    name: 'app',
    components: {
      HeaderMenu: HeaderMenu,
      Library: Library,
      SongDisplay: SongDisplay,
      ServiceOrder: ServiceOrder,
      Settings: Settings,
    },
    computed : mapState({
      settingsOn: state => state.settingsOn,
    }),
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
    #main,
    #library,
    #song-display {
      display: flex;
      flex-direction: column;
    }

    #header-menu,
    #app-body {
      display: flex;
      flex-direction: row;
      flex-basis: auto;
    }

    #main {
      height: 100%;
    }

    #header-menu {
      z-index: 200;
      height: 15%;
    }

    #app-body {
      z-index: 1;
      height: 85%;
    }
    
    #library {
      width: 25%;
      background: #ddd;
      animation-name: animatetop;
      animation-duration: 0.4s
    }

    #song-display {
      padding: 10px;
      flex-basis: 50%;
      overflow: auto;
    }

    #service-order {
      flex-basis: 25%;
      background: #ddd;
    }
  }

  @media (min-width: 690px) {
    
  }
</style>