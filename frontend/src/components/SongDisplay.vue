<template>
  <component :is="mode" id="display-wrapper"></component>
</template>

<script>
  import { mapState } from 'vuex'
  import { decodeSong, SongJSON } from '../decode'
  import SongReader from './SongReader.vue'
  import SongEditor from './SongEditor.vue'

  export default {
    name: "SongDisplay",
    components: {
      read: SongReader,
      edit: SongEditor,
    },
    computed: mapState({
      mode: state => state.mode,
    }),
  }
</script>

<style>
  #display-wrapper {
    flex-basis: 50%;
    padding: 10px;
    overflow: auto;
  }
</style>