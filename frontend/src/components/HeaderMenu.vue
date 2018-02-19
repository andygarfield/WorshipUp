<template>
  <header id="header-wrapper">
    <div id="header-content">
      <h1 id="mode-header">Plan</h1>
      <p>Service Date:</p>
      <Datepicker
        id="date-picker"
        :highlighted="highlighted"
        v-model="serviceDate"
        v-on:closed="toggleSelectingDate">
      </Datepicker>
    </div>
    <div id="mode-toggle">Present</div>
  </header>
</template>


<script>
  import Datepicker from 'vuejs-datepicker';
  import { mapState } from "vuex"

  export default {
    components: {
      Datepicker: Datepicker,
    },
    created () {
      this.getSetLists();
    },
    computed: {
      serviceDate: {
        get () {
          return this.$store.state.serviceDate
        },
        set (value) {
          this.$store.commit('setServiceDate', value)
        },
      },
      highlighted () {
        let parseDateString = function(dateString) {
          let year = dateString.slice(0, 4);
          let month = dateString.slice(4, 6) - 1;
          let day = dateString.slice(6, 8);
          return new Date(year, month, day);
        };
        return {
          dates: Object.keys(this.$store.state.setLists).map(el => parseDateString(el))
        }
      }
    },
    methods: {
      toggleSelectingDate () {
        this.selectingDate = !this.selectingDate
      },
      getSetLists () {
        return this.$store.dispatch("getSetLists")
      },
    },
  }
</script>

<style>
  #header-wrapper {
    display: flex;
    flex-direction: row;
    align-items: center;

    background: #bbb;
    padding: 0.5rem;
  }

  p {
    margin: 5px 0px;
  }

  h1 {
    margin: 0px 8px;
  }

  #header-content {
    display: flex;
    flex-direction: column;
    align-items: center;
    flex: 9;
  }

  #mode-toggle {
    flex: 1;
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