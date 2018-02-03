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
    #main {
      display: flex;
      flex-direction: column;
      height: 100%;
    }

    #header-menu {
      display: flex;
      flex-direction: row;
      height: 15%;
      flex-basis: auto;
    }

    #app-body {
      display: flex;
      flex-direction: row;
      height: 85%;
      flex-basis: auto;
    }
    
    #library {
      display: flex;
      flex-direction: column;
      width: 25%;
      background: #ddd;
    }

    #song-display {
      display: flex;
      flex-direction: column;
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