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

</style>