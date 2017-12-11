<template>
  <div>
    <h1>{{ songData.title }}</h1>
    <p><pre>{{ songData.lyrics }}</pre></p>
  </div>
</template>

<script>
  export default {
    name: "SongDisplay",
    props: ["songData"]
  }
</script>