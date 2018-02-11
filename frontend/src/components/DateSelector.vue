<template>
  <div id="main">
    <div id="month" class="column">
      <h2>Month</h2>
      <div class="digit">
        {{ month }}
        <div class="arrows">
          <div
            class="up"
            @click="columnUp('month')">
            ▲
          </div>
          <div
            class="down"
            @click="columnDown('month')">
            ▼
          </div>
        </div>
      </div>
    </div>
    <div id="day" class="column">
      <h2>Day</h2>
      <div class="digit">
        {{ day }}
        <div class="arrows">
          <div
            class="up"
            @click="columnUp('day')">
            ▲
          </div>
          <div
            class="down"
            @click="columnDown('day')">
            ▼
          </div>
        </div>
      </div>
    </div>
    <div id="year" class="column">
      <h2>Year</h2>
      <div class="digit">
        {{ year }}
        <div class="arrows">
          <div
            class="up"
            @click="columnUp('year')">
            ▲
          </div>
          <div
            class="down"
            @click="columnDown('year')">
            ▼
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "DateSelector",
  methods: {
    columnUp(column) {
      switch (column) {
        case "month":
          this.month++
          break;
      
        default:
          break;
      }
    }
  },
  data () {
    return {
      month: 1,
      day: 1,
      year: 2018,
    }
  }
}
</script>

<style scoped>
  #main {
    display: flex;
    flex-direction: row;
  }

  .column {
    display: flex;
    flex-direction: column;
    padding: 5px;
    text-align: center;
  }

  .digit {
    font-size: 24pt;
    display: flex;
    flex-direction: row;
    justify-content: center;
    border: 1px solid grey;
  }
  
  .arrows {
    font-size: 0.5em;
    display: flex;
    flex-direction: column;
  }
</style>
