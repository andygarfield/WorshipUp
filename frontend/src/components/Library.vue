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
          v-for="(songTitle, itemIndex) of songs"
          :key="itemIndex"
          @click="songClicked(songTitle)"
        >
          {{ songTitle }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  export default {
    name: "SongList",
    data() {
      return {
        songs: []
      }
    },
    props: ["refreshSongs"],
    created() {
      this.getSongList();
    },
    methods: {
      getSongList () {
        let xreq = new XMLHttpRequest();
        let vueInst = this;
        xreq.onload = function () {
          vueInst.songs = JSON.parse(this.responseText);
        }
        xreq.open("GET", "/songlist/", true);
        xreq.send();
      },

      songClicked (songTitle) {
        this.$emit("songClicked", songTitle);
      },

      newSong () {
        this.$emit("newSong");
      },
    },
    calculated: {
      refreshLibrary () {
        
      }
    }
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