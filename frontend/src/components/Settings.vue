<template>
  <div id="settings-background">
    <div
      v-click-outside="toggleSettings"
      id="settings">
      <span
        class="close"
        @click="toggleSettings">
          &times;
        </span>

      <h2>Import OpenSong songs:</h2>
      <form id="songupload" enctype="multipart/form-data" method="post">
        <input id="input-files" name="uploadfiles" type="file" multiple @change="submitSongs()">
      </form>

      <h2>Import OpenSong sets:</h2>
      <form id="setupload" enctype="multipart/form-data" method="post">
        <input id="input-files" name="uploadfiles" type="file" multiple @change="submitSets()">
      </form>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  
  methods: {
    toggleSettings () {
      this.$store.commit("toggleSettings");
    },
    submitSongs () {
      var form = document.forms.namedItem("songupload");
      let fData = new FormData(form);

      let xhr = new XMLHttpRequest();
      xhr.open("POST", "/songupload/" , true);
      xhr.onload = (res) => {
        this.$store.dispatch("getSongList");
        this.$store.commit("toggleSettings");
      }

      xhr.send(fData)
    },
    submitSets () {
      var form = document.forms.namedItem("setupload");
      let fData = new FormData(form);

      let xhr = new XMLHttpRequest();
      xhr.open("POST", "/setupload/" , true);
      xhr.onload = (res) => {
        this.$store.dispatch("getSetLists");
        this.$store.commit("toggleSettings");
      }

      xhr.send(fData)
    }
  },
  directives: {
    // Stolen with love from https://jsfiddle.net/Linusborg/Lx49LaL8/
    'click-outside': {
      bind: function(el, binding, vNode) {
        // Provided expression must evaluate to a function.
        if (typeof binding.value !== 'function') {
        	const compName = vNode.context.name
          let warn = `[Vue-click-outside:] provided expression '${binding.expression}' is not a function, but has to be`
          if (compName) { warn += `Found in component '${compName}'` }
          
          console.warn(warn)
        }
        // Define Handler and cache it on the element
        const bubble = binding.modifiers.bubble
        const handler = (e) => {
          if (bubble || (!el.contains(e.target) && el !== e.target)) {
          	binding.value(e)
          }
        }
        el.__vueClickOutside__ = handler

        // add Event Listeners
        document.addEventListener('click', handler)
			},
      
      unbind: function(el, binding) {
        // Remove Event Listeners
        document.removeEventListener('click', el.__vueClickOutside__)
        el.__vueClickOutside__ = null

      }
    }
  },
}
</script>

<style scoped>
  #settings-background {
    display: flex;
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
    z-index: 2;
    display: flex;
    flex-direction: column;
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
