<template>
  <div v-html="songHtml"></div>
</template>

<script>
  export default {
    name: "SongDisplay",
    props: ["songHtml"]
  }
</script>

<style>
  h1 {
    margin-top: 0px;
  }
</style>