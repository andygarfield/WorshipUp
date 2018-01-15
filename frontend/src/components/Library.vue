<template>
  <div>
    <div id="library-top-menu">
      <div
      id="add-button"
      @click="newSong">
        Add New
      </div>
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
      }
    },
  }
</script>

<style>

#add-button {
  height: 30px;
  display: flex;
  background: lightgreen;
  justify-content: center;
  align-items: center;
}

</style>