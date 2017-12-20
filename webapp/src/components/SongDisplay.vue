<template>
  <div>
    <div v-html="songHtml" v-if="mode == 'read'"></div>
    <form
      id="edit-song"
      v-if="mode == 'edit'"
      action="/newSong">
      <button
        type="submit"
        form="edit-song">
        Save
      </button>
      <input type="text">
    </form>
  </div>
</template>

<script>
  export default {
    name: "SongDisplay",
    props: ["songHtml", "mode"]
  }
</script>

<style>
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
</style>