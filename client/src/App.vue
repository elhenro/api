<template>
  <div id="app">
    <p>{{results}}</p>
     <vue-circle
        :progress=results
        :size="100"
        :reverse="false"
        line-cap="round"
        :fill="fill"
        empty-fill="rgba(0, 0, 0, .1)"
        :animation-start-value="0.0"
        :start-angle="0"
        insert-mode="append"
        :thickness="5"
        :show-percent="true"
        @vue-circle-progress="progress"
        @vue-circle-end="progress_end">
          <p>Slot!</p>
      </vue-circle>
  </div>
</template>

<script>
  import Vue from 'vue'
  import axios from 'axios'
  import VueAxios from 'vue-axios'
  import VueCircle from 'vue2-circle-progress'
  
  Vue.use(VueAxios, axios)

const url = "http://localhost:8000/time/percent";

export default {
  name: 'app',
  components: {
    VueCircle
  },
  data () {
    return {
      results: [],
      msg: 'Welcome to Your Vue.js App',
      fill : { gradient: ["red", "green", "blue"] },
    }
  },
  methods:{
    progress(event,progress,stepValue){
      console.log(stepValue);
    },
    progress_end(event){
      console.log("Circle progress end");
    }
  },
  mounted() {
    axios.get(url).then(response => {
      this.results = response.data
    })
  },
  render: h => h(App)
}
</script>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
