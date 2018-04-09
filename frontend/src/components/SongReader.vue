<template>
  <div>
    <h1 v-if="notBlank" id="song-title">{{ song.title }}</h1>
    <div id="action-icons">
      <img
        v-if="notBlank"
        @click="addToService(song.title)"
        src="/static/plus.svg"
        class="modify-btn servadd-btn">
      <img
        v-if="notBlank"
        @click="switchMode('edit')"
        src="/static/edit.svg"
        class="modify-btn edit-btn">
      <img
        v-if="notBlank"
        @click="deleteSong"
        src="/static/delete.svg"
        class="modify-btn delete-btn">
    </div>
    <div
      id="song-body"
      v-html="songBody"
      v-if="notBlank">
    </div>
  </div>
  
</template>

<script>
  import Vue from 'vue'
  import { mapState } from 'vuex'
  import { decodeSong } from '../decode'

  export default Vue.extend({
    name: "SongReader",
    computed: mapState({
      song: state => state.displayedSong.song,
      songBody: state => decodeSong(state.displayedSong.song),
      notBlank(state) {
        return state.displayedSong.i != -1;
      },
    }),
    methods: {
      switchMode (newMode) {
        this.$store.commit("switchMode", "edit");
      },
      addToService () {
        this.$store.commit("addToService", this.song.title);
      },
      deleteSong (id) {
        this.$store.dispatch("deleteSong", this.song.id);
      },
    }
  })
</script>

<style>
  #song-title {
    margin: 0 0 5px;
  }

  #action-icons {
    display: flex;
  }

  #song-body {
    font-size: 12pt;
  }

  .couplet-line {
    display: flex;
    flex-flow: row wrap;
  }

  .chord {
    color: palevioletred;
  }

  .song-section {
    font-weight: bold;
    margin: 15px 0px 5px;
    font-size: 1.1em;
  }

  .modify-btn {
    cursor: pointer;
    height: 22px;
    width: 22px;
    padding: 2px;
    border: 2px solid;
    border-radius: 3px;
    margin-right: 3px;
  }

  .edit-btn {
    border-color: #ff9900;
  }

  .delete-btn {
    border-color: #a8050f;
  }

  .servadd-btn {
    border-color: #0087ee;
  }

  .comment {
    color: dimgrey;
    font-style: italic;
  }
</style>