<template>
  <section id="library-wrapper">
    <div id="library-top-menu">
      <div
      id="add-button"
      @click="newSong">
          Add New
      </div>
      <img
      src="/static/settings.svg"
      id="settings-button"
      @click="toggleSettings()" />
    </div> 
    <div id="library-elements">
      <div
        class="list-element"
        v-for="(item, sKey) of songList"
        :key="sKey"
        @click="getSong(item)"
      >
        {{ item.song.title }}
      </div>
    </div>
  </section>
</template>

<script>
  import { mapState } from "vuex"
  import { IndexedSong } from "../main"

  export default {
    name: "SongList",
    computed: {
      songList ()  {
        return this.$store.state.songList;
      },
    },
    created () {
      return this.$store.dispatch("getSongList");
    },
    methods: {
      getSong (songItem) {
        return this.$store.dispatch("getSong", songItem);
      },
      // changeIndex (songIdx) {
      //   return this.$store.commit("changeSongIndex", songIdx)
      // },
      newSong () {
        this.$store.commit("makeSongBlank");
        return this.$store.commit("switchMode", "edit");
      },
      toggleSettings () {
        return this.$store.commit("toggleSettings");
      },
    },
  }
</script>

<style>
#library-wrapper {
  flex-basis: 25%;

  display: flex;
  flex-direction: column;

  background: #ddd;
}

#library-top-menu {
  display: flex;
  flex-direction: row;
}

#add-button {
  display: flex;
  flex-grow: 1;
  height: 30px;
  background: lightgreen;
  justify-content: center;
  align-items: center;
}

#settings-button {
  height: 30px;
  width: 30px;
  padding: 3px;
  background: #929292;
}

#library-elements {
  flex: 1;

  display: flex;
  flex-direction: column;
  overflow: auto;
  margin-top: 5px;
}

.list-element {
  padding: 0 5px;
}

</style>