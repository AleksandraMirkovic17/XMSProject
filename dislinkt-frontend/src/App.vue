<template>
  <div id="app">
    <router-view name="header" />

    <div class="wrapper">
      <router-view />
    </div>
    <router-view name="footer" />
  </div>
</template>

<script>

import router from "./router";

export default {
  name: 'App',
  data(){
    return {
      user: {},
      searchParams: ''
    }
  },
  components: {

  },
  mounted: function() {
      this.setupData()
    console.log("here")
  },
  computed: {
    getUserProfileHref() {
      return "/profile/" + (JSON.parse(this.user)).username
    },
  },
  methods: {
    setupData() {
      this.user = localStorage.getItem('user')
    },
    logout() {
      localStorage.removeItem('user');
    },
    searchProfiles(){
      if(this.searchParams!=''){
        router.push('/profiles/'+this.searchParams)
      }
    }
  },
  watch: {
    $route () {
      this.setupData()
    }
  }
}
</script>

<style>


</style>
