<template>
  <div id="settings-background">
    <div id="settings">
      <span
        class="close"
        @click="toggleSettings"
        >
          &times;
        </span>
      <h2>Import songs:</h2>
      <input id="input-files" type="file" multiple>
      <input @click="submitSongs" type="submit">
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  methods: {
    toggleSettings () {
      this.$store.commit("toggleSettings")
    },
    submitSongs () {
      console.log(document.getElementById("input-files"))
    }
  }
}
</script>

<style>
  #settings-background {
    position: fixed;
    z-index: 1;
    padding-top: 100px;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: auto; /* Enable scroll if needed */
    background-color: rgb(0,0,0); /* Fallback color */
    background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
  }
  
  #settings {
    background: #fefefe;
    margin: 15% auto;
    padding: 20px;
    border: 1px solid #888;
    width: 40%;
  }

.close {
    cursor: pointer;
    color: #aaa;
    float: right;
    font-size: 40px;
    font-weight: bold;
}
</style>
