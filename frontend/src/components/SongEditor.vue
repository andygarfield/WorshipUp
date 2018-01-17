<template>
  <div id="edit-song">
    <div id="form-area">
      <button id="save-button" type="submit" form="edit-song" @click="save">
        Save
      </button>
      <input id="edit-title" placeholder="Title" :value="songTitle">
      <textarea id="edit-body" placeholder="Song Body" :value="songBody"></textarea>
    </div>
  </div>
</template>

<script>
  export default {
    computed: {
      songTitle () {
        return this.$store.state.songData.title;
      },
      songBody () {
        return this.$store.state.songData.body;
      }
    },
    methods: {
      save () {
        let songTitle = document.getElementById("edit-title");
        let songBody = document.getElementById("edit-body");

        let newSongBody = encodeURI(songBody.value)

        let xhttp = new XMLHttpRequest();
        
        xhttp.open("POST", "/songsubmit");
        xhttp.onload = () => {
          if (xhttp.response.substr(0, 5) != "Error") {
            this.$store.dispatch("getSongList");
            this.$store.dispatch("loadSong", songTitle.value);
            this.$store.commit("switchMode", "read");
          } else {
            alert(xhttp.response);
          }
        }

        xhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xhttp.send("title=" + songTitle.value  + "&" + "body=" + newSongBody);
      },
      switchMode: newMode => state.commit("switchMode", newMode)
    },
  }
</script>

<style>
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

  input#edit-title {
    border: none;
    background: #dadada;
    font-size: 1.5em;
    padding: 8px;
  }

  textarea#edit-body {
    width: 100%;
    height: 100%;
    padding: 8px;
    box-sizing: border-box;
    border: none;
    font-family: Consolas, Monaco, Lucida Console, Liberation Mono, DejaVu Sans Mono, Bitstream Vera Sans Mono, Courier New, monospace;
  }
</style>