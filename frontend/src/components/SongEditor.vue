<template>
  <div id="edit-song">
    <div id="form-area">
      <button id="save-song" type="submit" form="edit-song" @click="save">
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
      currentIndex () {
        return this.$store.state.displayedSong.i;
      },
      songTitle () {
        return this.$store.state.displayedSong.song.title;
      },
      songBody () {
        return this.$store.state.displayedSong.song.body;
      }
    },
    methods: {
      save () {
        let songTitle = document.getElementById("edit-title").value;
        let songBody = document.getElementById("edit-body").value;

        // Replacing new lines with ~ because github.com/graph-gophers/graphql-go
        // is currently using go's text/scanner, which doesn't handle new lines
        // correctly
        if (this.currentIndex == -1) {
          this.$store.dispatch("addSong", {
            title: songTitle,
            body: songBody.replace(/\n/g, "~"),
          })
        } else {
          this.$store.dispatch("updateSong", {
            id: this.$store.state.displayedSong.id,
            title: songTitle,
            body: songBody.replace(/\n/g, "~"),
          })
        }


        // let newSongBody = encodeURIComponent(songBody.value)

        // let xhttp = new XMLHttpRequest();
        
        // xhttp.open("POST", "/songsubmit/");
        // xhttp.onload = () => {
        //   if (xhttp.response.substr(0, 5) != "Error") {
        //     this.$store.dispatch("allSongs");
        //     this.$store.dispatch("getSongs", songTitle.value);
        //     this.$store.commit("switchMode", "read");
        //   } else {
        //     alert(xhttp.response);
        //   }
        // }

        // xhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        // xhttp.send("title=" + songTitle.value  + "&" + "body=" + newSongBody);
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

  #save-song {
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
    font-size: 10pt;
    box-sizing: border-box;
    border: none;
    font-family: Consolas, Monaco, Lucida Console, Liberation Mono, DejaVu Sans Mono, Bitstream Vera Sans Mono, Courier New, monospace;
  }
</style>