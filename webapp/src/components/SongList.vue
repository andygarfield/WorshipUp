<template>
  <div>
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
</template>

<script>
  export default {
    name: "SongList",
    data() {
      return {
        songs: []
      }
    },
    created() {
      let xreq = new XMLHttpRequest();
      let vueInst = this;
      xreq.onreadystatechange = function () {
        if (this.readyState == 4 && this.status == 200) {
          vueInst.songs = JSON.parse(this.responseText);
        }
      }
      xreq.open("GET", "/songlist/", true);
      xreq.send();
    },
    methods: {
      songClicked (songTitle) {
        this.$emit("songClicked", songTitle)
      }
    }
  }
</script>

<style>

</style>