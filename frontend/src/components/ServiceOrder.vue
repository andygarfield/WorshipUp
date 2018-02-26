<template>
  <div id="service-wrapper">
    <draggable v-model="setSongs">
      <transition-group>
        <div
          v-for="(song, songIndex) in setSongs"
          :key="songIndex"
          class="button set-item">
          {{ song }}
        </div>
      </transition-group>
    </draggable>
    <div
      id="save-set"
      class="button"
      v-if="modified"
      @click="saveSet">
        Save
    </div>
  </div>
</template>

<script>
  import Vue from 'vue'
  import draggable from 'vuedraggable'

  export default {
    components: {
      draggable,
    },
    computed: {
      setSongs: {
        get () {
          return this.$store.getters.setSongs;
        },
        set (value) {
          this.modified = true;
          let updatedSet = {
            stringDate: this.stringDate,
            setList: value,
          }
          this.$store.commit('updateSetList', updatedSet)
        }
      },
      stringDate () {
        return this.$store.getters.stringDate;
      }
    },
    data () {
      return {
        modified: false
      }
    },
    methods: {
      saveSet () {
        let xhttp = new XMLHttpRequest();
        
        xhttp.open("POST", "/setsubmit/");
        xhttp.onload = () => {
          if (xhttp.response.substr(0, 5) != "Error") {
            this.$store.dispatch('getSetLists')
            this.modified = false;
          } else {
            alert(xhttp.response)
          }
        }

        xhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        xhttp.send("date=" + this.stringDate  + "&" + "set=" + encodeURIComponent(this.setSongs));
      },
    }
  }
</script>

<style>
  #service-wrapper {
    display: flex;
    flex-direction: column;
    align-items: stretch;
    flex-basis: 25%;
    background: #ddd;
    overflow: auto;
  }

  .button {
    display: flex;
    justify-content: center;
    align-items: center;

    cursor: pointer;

    color: white;
    border-radius: 100px;
  }

  .set-item {
    background: rgb(145, 145, 145);
    padding: 4px 4px 7px;
    margin: 5px;
    box-shadow: 0 2px 0 0 rgba(84, 86, 88, .06);
  }

  #save-set {
    background-color: #007bff;
    border-color: #007bff;
    box-shadow: 0 2px 0 0 rgba(65, 160, 255, 0.603);
    border-width: 1px;
    align-self: center;
    padding: 8px 11px;
    margin: 10px;
    width: fit-content;
  }
</style>
