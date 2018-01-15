<template>
  <div>
    <h1 id="song-title">{{ songTitle }}</h1>
    <img
      v-if="songBody"
      @click="switchMode('edit')"
      src="/static/edit.svg"
      class="modify-btn edit-btn">
    <img
      v-if="songBody"
      src="/static/delete.svg"
      class="modify-btn delete-btn">
    <div id="song-body" v-html="songBody"></div>
  </div>
  
</template>

<script>
  import { mapState } from 'vuex'
  import { decodeSong } from '../decode'

  export default {
    name: "SongReader",
    computed: mapState({
      songTitle: state => {
        return state.songData.title
      },
      songBody: state => {
        return decodeSong(state.songData)
      }
    }),
    methods: {
      switchMode (newMode) {
        this.$store.commit("switchMode", "edit")
      }
    }
  }
</script>

<style>
  h1 {
    margin-top: 0px;
    margin-bottom: 5px;
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
    height: 15px;
    width: 15px;
    padding: 2px;
    border: 2px solid;
    border-radius: 3px;
  }

  .edit-btn {
    border-color: black;
  }

  .delete-btn {
    border-color: #a8050f;
  }

  .comment {
    color: dimgrey;
    font-style: italic;
  }
</style>