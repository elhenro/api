<template>
  <div id="app">

    <input v-model="time" placeholder=".. time ..." v-on:keyup.enter="sendTime()">
    <button v-on:click="sendTime()">go</button>
    <p>{{results}}</p>

    <br>
    <vue-circle
      ref="twheel1"
      :progress=results
      :size="200"
      :reverse="false"
      line-cap="round"
      :fill="fill"
      empty-fill="rgba(0, 30, 0, .1)"
      :animation-start-value="0.0"
      :start-angle="0"
      insert-mode="append"
      :thickness="5"
      :show-percent="true"
      @vue-circle-progress="progressP"
      @vue-circle-end="progress_end">
        <p>Slot!</p>
      </vue-circle>
      
      <div class="chatWindow">
        <p v-for="msg in messages">
          {{msg.id}} {{msg.name}} {{msg.Creation}}
        </p>
      </div>

      <input v-model="message" placeholder="..msg ..." v-on:keyup.enter="sendMessage()">
      <button v-on:click="sendMessage()">go</button>
  </div>
</template>

<script>
  import Vue from 'vue'
  import TypeDetect from 'type-detect'
  import axios from 'axios'
  import VueAxios from 'vue-axios'
  import VueCircle from 'vue2-circle-progress-redraw'

  
  Vue.use(VueAxios, axios)

const url = "http://localhost:8000/time/percent";
const urlReadTextApi = "http://localhost:8000/text";
const urlWriteTextApi = "http://localhost:8000/text/"/*parameter*/;

export default {
  name: 'app',
  components: {
    VueCircle
  },
  data () {
    return {
      results: 0,
      time: "",
      message: "",
      messages: Array,/*[]
        { text: "tt"},
        { text: "tddt"}
      ],*/
      fill: { gradient: ["black", "red"]},
    }
  },
  methods:{
    debug(boolean){
      return false;
    },
    progressP(event,progress,stepValue){
      if (this.debug()){
        console.log(stepValue);
      }
    },
    progress_end(event){
      if (this.debug()){
        console.log("Circle progress end");
      }
    },
    setTime(string){
      //timeValueString = "10:30/18:01"
      var reqUrl = 'http://localhost:8000/time/set/' + this.getTime();
      axios.post(reqUrl, 
      this.name,
      { headers: {
        'Content-type': 'application/x-www-form-urlencoded',
      }
      })/*.then(response => ); */
      this.refreshTime(this.getProgress())
    },
    sendTime(){
      this.setTime();
    },
    getTime(){
      return this.time;
    },
    sendMessage(){
      this.addMessage(this.getMessage)
      //this.requestMessages()
      this.sleep(200)
      this.loadMessages()
      this.sleep(200)
      this.loadMessages()
      this.message = ""
    },
    sleep(ms){
       return new Promise(resolve => setTimeout(resolve, ms))
    },
    getMessage(){
      // gets message user input
      return this.message;
    },
    returnTest(){
      return "test"
    },
    loadMessages() {
      var r = Array
      axios.get(urlReadTextApi).then(response => {
        this.setLocalMessages(response.data)
      })
    },
    setLocalMessages(array){
      this.messages = array.reverse()
    },
    addMessage(string){
      axios.post(urlWriteTextApi + this.getMessage(), 
      this.name,
      { headers: {
          'Content-type': 'application/x-www-form-urlencoded',
        }
      })
    },
    refreshTime(){
      axios.get(url).then(response => {
        this.results = Number(response.data)
      })

      this.$refs.twheel1.updateProgress(parseInt(this.results))
      this.$refs.twheel1.redraw({ animation: { duration: 5000 }, size: 150, fill: this.fill });
    },
    checktype(obj, type) {
      return isType(obj === type)
    },
    getProgress(){
      var content = Number
      axios.get(url).then(response => {
        //content = Math.imul(10 * Number(response.data))
        content = response.data
      })
      return content
    }
  },
  mounted() {
    this.refreshTime()
    this.loadMessages()
    //this.messages = this.getMessages()
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
.chatWindow {
  height: 40vh;
  overflow: auto;
}
</style>
