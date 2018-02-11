<template>
  <div>
    <div id="header-content">
      <h1 id="mode-header">Plan</h1>
      <p>Service Date</p>
      <Datepicker
        id="date-picker"
        v-model="serviceDate"
        v-on:closed="toggleSelectingDate">
      </Datepicker>
    </div>
    <div id="mode-toggle">Present</div>

    <!-- <div
      id="datesetter-background"
      v-if="selectingDate">
      <div
        id="datesetter">
        <Datepicker v-model="serviceDate" v-on:closed="toggleSelectingDate"></Datepicker>
      </div>
    </div> -->

  </div>
</template>

<script>
  import Datepicker from 'vuejs-datepicker';
  import { mapState } from "vuex"

  export default {
    components: {
      Datepicker: Datepicker,
    },
    computed: {
      serviceDate: {
        get () {
          return this.$store.state.serviceDate
        },
        set (value) {
          this.$store.commit('setServiceDate', value)
        },
      }
    },
    methods: {
      toggleSelectingDate () {
        this.selectingDate = !this.selectingDate
      }
    },
  }
</script>

<style scoped>
  p {
    margin: 5px 0px;
  }

  #header-menu {
    display: flex;
    flex-flow: row nowrap;
    justify-content: space-around;
    align-items: stretch;
    overflow: hidden;
    background: #bbb;
  }

  #header-content {
    display: flex;
    flex-direction: column;
    align-items: center;
    flex-grow: 9;
  }

  #mode-toggle {
    flex-grow: 1;
    text-align: center;
  }

  #date-picker {
    cursor: pointer;
    border: 1px solid grey;
    padding: 4px;
    width: fit-content;
  }

.close {
    cursor: pointer;
    color: #aaa;
    float: right;
    font-size: 40px;
    font-weight: bold;
}
</style>