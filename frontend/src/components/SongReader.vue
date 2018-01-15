<template>
  <div id="song-read" v-html="songData">
  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import { decodeSong } from '../decode'

  export default {
    name: "SongReader",
    computed: mapState({
      songData: state => {
        return decodeSong(state.songData)
      }
    })
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
</style>