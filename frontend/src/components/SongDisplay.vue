<template>
  <div>
    <div
      id="song-read"
      v-html="songHtml"
      v-if="mode == 'read'">
    </div>
    <div
      id="edit-song"
      v-if="mode == 'edit'"
      >
      <button
        id="save-button"
        type="submit"
        form="edit-song"
        @click="save">
        Save
      </button>
      <textarea id="edit-text"></textarea>
    </div>
  </div>
</template>

<script>
  export default {
    name: "SongDisplay",
    props: ["songHtml", "mode"],
    methods: {
      save () {
        let xhttp = new XMLHttpRequest()
        xhttp.open("POST", "/newSong")
        xhttp.onload = function(e) {
          console.log(e);
        }

        let text = document.getElementById("edit-text");
        
        xhttp.send(text.innerText)
      }
    }
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

  #edit-song {
    display: flex;
    flex-flow: column nowrap;
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

  textarea#edit-text {
    width: 100%;
    height: 100%;
    padding: 17px;
    box-sizing: border-box;
    border: none;
    font-family: Consolas, Monaco, Lucida Console, Liberation Mono, DejaVu Sans Mono, Bitstream Vera Sans Mono, Courier New, monospace;
  }

</style>