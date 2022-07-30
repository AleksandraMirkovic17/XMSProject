<template>
  <nav class = "navbar navbar-expand-lg bg-dark navbar-dark mainavbar" style="height: 8%">
    <div class="container">
      <a v-if="user === null" href="/home" class="navbar-brand">Adventureland</a>
    </div>
  </nav>
    <router-view/>
</template>

<script>
//import {useStore} from "vuex"
//import axios from 'axios'
//import axios from "axios";
//import {devServer} from "../vue.config";
export default{
  data(){
    return{
      user: null,
      loggedUser: null
    }
  },
  watch: {

}
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
.mainavbar{
  position: fixed;
  z-index: 10;
  width: 100%;
  margin-bottom: 1%;
}
</style>