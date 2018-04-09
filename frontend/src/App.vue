<template>
  <div id="app">
    <Settings v-if="settingsOn"></Settings>
    <HeaderMenu>
    </HeaderMenu>
    <main id="main-wrapper">
      <Library
        @newSong="newSong">
      </Library>
      <SongDisplay></SongDisplay>
      <ServiceOrder></ServiceOrder>
    </main>
  </div>
</template>

<script>
  import HeaderMenu from './components/HeaderMenu.vue';
  import Library from './components/Library.vue';
  import SongDisplay from './components/SongDisplay.vue';
  import ServiceOrder from './components/ServiceOrder.vue';
  import Settings from  './components/Settings.vue';
  import decodeSong from './decode.ts';
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
  /* @media (min-width: 0px) { */
    #app {
      display: flex;
      flex-direction: column;
    }

    #main-wrapper {
      min-height:0;
      display: flex;
      flex: 1;
      flex-direction: row;
    }

    #service-order {
      flex-basis: 25%;
      background: #ddd;
    }
  /* } */

  /* @media (min-width: 690px) {
    
  } */
</style>