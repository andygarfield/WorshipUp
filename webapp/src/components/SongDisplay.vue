<template>
  <div>
    <div v-html="songHtml" v-if="mode == 'read'"></div>
    <input v-if="mode == 'edit'">
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
</style>