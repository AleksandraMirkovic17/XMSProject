<template>
  <nav class = "navbar navbar-expand-lg bg-dark navbar-dark mainavbar" style="height: 7%; position: fixed;">
    <div class="container">
      <a  class="navbar-brand" href="/home">Agent application</a>
      <!--
      <div class="collapse navbar-collapse">
        <ul class ="navbar-nav ms-auto">
          <li class="nav-item">
            <a class="nav-link" href="/home" v-on:click="logOut">Log out</a>

          </li>
          <li class="nav-item">
            <a class="nav-link" href="/registrationCompany" >Company registration</a>

          </li>

           <li class="nav-item">
            <a class="nav-link" href="/confirmCompany" >All companies (ADMIN SIDE)</a>

          </li>
          <li class="nav-item">
            <a class="nav-link" href="/listingCompany" >Comapny listing (USER SIDE)</a>

          </li>
        </ul>
      </div>-->
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
 
    }
  },
  watch: {

},
  
  mounted(){

  },
  methods:{
    logOut(){
      let router = this.$router;
      localStorage.removeItem('token');
      localStorage.clear();
      router.push("/login");
      return;
    }
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