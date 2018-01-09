<template>
  <div>
    <div
      id="song-read"
      v-html="songData"
      v-if="mode == 'read'">
    </div>
    <div
      id="edit-song"
      v-if="mode == 'edit'"
      >
      <div id="form-area">
        <button
          id="save-button"
          type="submit"
          form="edit-song"
          @click="save">
          Save
        </button>
        <textarea id="edit-title"></textarea>
        <textarea id="edit-body"></textarea>
      </div>
    </div>
  </div>
</template>

<script>
  import { mapState } from 'vuex'

  export default {
    name: "SongDisplay",
    methods: {
      save () {
        let vueInst = this;
        let xhttp = new XMLHttpRequest();
        
        xhttp.open("POST", "/newSong");
        xhttp.onload = function(e) {
          vueInst.$store.dispatch("getSongList");
        }
        let songTitle = document.getElementById("edit-title");
        let songBody = document.getElementById("edit-body");

        xhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xhttp.send("title=" + songTitle.value  + "&" + "body=" + songBody.value);
      },
    },
    computed: mapState({
      mode: state => state.mode,
      songData: state => state.songData,
    }),
  }
</script>

<style>
  #song-read {
    padding: 15px;
  }

  h1 {
    margin-top: 0px;
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

  #form-area {
    display: flex;
    flex-flow: column nowrap;
    box-shadow: 1px 1px 3px 0px rgba(0,0,0,0.3);
    height: 100%;
    box-sizing: border-box;
    border-radius: 4px 4px 0 0;
    min-height: 100%;
    min-width: 100%;
  }

  #edit-song {
    background: #eee;
    height: 100%;
    padding: 20px;
    box-sizing: border-box;
  }

  #save-button {
    background: lightgreen;
    border-radius: 4px 4px 0 0;
    padding: 10px;
    border: none;
    cursor: pointer;
  }

  textarea#edit-body {
    width: 100%;
    height: 100%;
    padding: 17px;
    box-sizing: border-box;
    border: none;
    font-family: Consolas, Monaco, Lucida Console, Liberation Mono, DejaVu Sans Mono, Bitstream Vera Sans Mono, Courier New, monospace;
  }

</style>