<template>
  <div id="edit-song">
    <div id="form-area">
      <button id="save-button" type="submit" form="edit-song" @click="save">
        Save
      </button>
      <input id="edit-title" placeholder="Title"></input>
      <textarea id="edit-body" placeholder="Song Body"></textarea>
    </div>
  </div>
</template>

<script>
  export default {
    methods: {
      save () {
        let songTitle = document.getElementById("edit-title");
        let songBody = document.getElementById("edit-body");

        let xhttp = new XMLHttpRequest();
        
        xhttp.open("POST", "/newSong");
        xhttp.onload = () => {
          if (xhttp.response != "Error: Invalid input") {
            this.$store.dispatch("getSongList");
            this.$store.dispatch("loadSong", songTitle.value);
            this.$store.commit("switchMode", "read");
          } else {
            alert("Error: Invalid input")
          }
        }

        xhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xhttp.send("title=" + songTitle.value  + "&" + "body=" + songBody.value);
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