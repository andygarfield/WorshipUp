<template>
  <div>
    <div id="header-content">
      <h1 id="mode-header">Plan</h1>
      <h2>Service Date</h2>
    </div>
    <div id="mode-toggle">Present</div>
  </div>
</template>

<script>
  export default {

  }
</script>

<style>
  #header-menu {
    display: flex;
    flex-flow: row nowrap;
    justify-content: space-around;
    align-items: stretch;
    overflow: hidden;
    background: #bbb;
  }

  #header-content {
    justify-self: center;
    text-align: center;
    flex-grow: 9;
  }

  #mode-toggle {
    flex-grow: 1;
    text-align: center;
  }
</style>