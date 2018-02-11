<template>
  <div id="library">
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
      <ul>
        <li
          v-for="(songTitle, itemIndex) of songList"
          :key="itemIndex"
          @click="loadSong(songTitle)"
        >
          {{ songTitle }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  import { mapState } from "vuex"

  export default {
    name: "SongList",
    computed: mapState({
      songList: state => state.songList,
    }),
    created () {
      this.getSongList();
    },
    methods: {
      loadSong (songTitle) {
        return this.$store.dispatch("loadSong", songTitle)
      },
      newSong () {
        this.$store.commit("changeSongData", `{"title": "", "body": ""}`)
        return this.$store.commit("switchMode", "edit")
      },
      getSongList () {
        return this.$store.dispatch("getSongList")
      },
      toggleSettings () {
        return this.$store.commit("toggleSettings")
      },
    },
  }
</script>

<style>

#library-top-menu {
  display: flex;
}

#library-elements {
  display: flex;
  height: 100%;
  overflow: auto;
}

#library-elements li {
  cursor: pointer;
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
  flex-basis: 24px;
  height: 24px;
  padding: 3px;
  background: #929292;
}

</style>