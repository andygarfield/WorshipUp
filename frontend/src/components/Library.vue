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
        v-for="(songTitle, itemIndex) of songList"
        :key="itemIndex"
        @click="loadSong(songTitle)"
      >
        {{ songTitle }}
      </div>
    </div>
  </section>
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

  padding: 5px;
}

</style>