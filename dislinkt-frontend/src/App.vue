<template>
  <div id="app">
    <nav class="navMenu navbar-dark bg-dark" style="padding: 1%">
      <a href="/" style="float:left">Homepage</a>
      <a href="/login">Login</a>
      <a href="/register">Join us</a>
      <a>
        <input class="form-control" list="datalistOptions" id="exampleDataList" v-model="searchParams" placeholder="Type to search...">
      </a>
      <a>
        <button class="btn-primary" v-on:click="searchProfiles()">Search</button>
      </a>

    </nav>

    <router-view/>
  </div>
</template>

<script>

import router from "./router";

export default {
  name: 'App',
  data(){
    return {
      searchParams: ''
    }
  },
  components: {

  },
  methods: {
    searchProfiles(){
      if(this.searchParams!=''){
        router.push('/profiles/'+this.searchParams)
      }
    }
  },
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Roboto:wght@100;300&display=swap');

#app {
  font-family: 'Roboto', sans-serif;
  font-size: 12px;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
}

body, html {
  height:100vh;
  min-height:550px;
  margin: 0;
  background-color: #424242;
  background-position: center;
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
  font-family: 'Roboto', sans-serif;
}

input {
    outline: none;
    border: none;
}

.navMenu {
  align-content: right;
  padding-left:20px !important;
  padding-right:20px !important;
  padding-top:3px !important;
  text-align: right;
  width: 100%;
  font-size: 16px;
}

.navMenu a {
  margin-left: 10px;
  margin-right: 10px;
}

.navMenu a {
	color: #eae7dc;
	text-decoration: none;
	text-transform: uppercase;
	font-weight: 500;
	display: inline-block;
	-webkit-transition: all 0.2s ease-in-out;
	transition: all 0.2s ease-in-out;
}

.navMenu a:hover {
  color: #e85a4f
}
</style>
