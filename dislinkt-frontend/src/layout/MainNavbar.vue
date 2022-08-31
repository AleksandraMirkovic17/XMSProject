<template>
  <navbar
    position="fixed"
    type="primary"
    :transparent="transparent"
    :color-on-scroll="colorOnScroll"
    menu-classes="ml-auto"
  >
      <template >
        <router-link class="navbar-brand" to="/">
          <p href="/" style="font-size: 20pt">Dislinkt</p>
        </router-link>
      </template>
    <template class="justify-content-end" slot="navbar-menu">
      <li class="nav-item" v-if="user">
        <a
          class="nav-link"
          :href="getUserProfileHref"
          target="_blank"
        >
          <i class="now-ui-icons users_single-02"></i>
          <p>Profile</p>
        </a>
      </li>
      <drop-down
              tag="li"
              title="Join us"
              icon="now-ui-icons emoticons_satisfied"
              class="nav-item"
              v-if="!user"
      >
        <nav-link to="/login">
          <i class="now-ui-icons users_circle-08"></i> Login
        </nav-link>
        <nav-link to="/register">
          <i class="now-ui-icons users_single-02"></i> Register
        </nav-link>
      </drop-down>
      <li class="nav-item" v-if="user" v-on:click="logout">

                <a class="nav-link btn btn-neutral">
                  <div  style="display: flex; flex-direction: row; ">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-box-arrow-in-right" viewBox="0 0 16 16">
                      <path fill-rule="evenodd" d="M6 3.5a.5.5 0 0 1 .5-.5h8a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5h-8a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 0-1 0v2A1.5 1.5 0 0 0 6.5 14h8a1.5 1.5 0 0 0 1.5-1.5v-9A1.5 1.5 0 0 0 14.5 2h-8A1.5 1.5 0 0 0 5 3.5v2a.5.5 0 0 0 1 0v-2z"/>
                      <path fill-rule="evenodd" d="M11.854 8.354a.5.5 0 0 0 0-.708l-3-3a.5.5 0 1 0-.708.708L10.293 7.5H1.5a.5.5 0 0 0 0 1h8.793l-2.147 2.146a.5.5 0 0 0 .708.708l3-3z"/>
                    </svg>
                    <p style="padding-left: 5pt">Logout</p>
                  </div>


                </a>

      </li>
      <li class="nav-item">
        <div style="display: flex; flex-direction: row; margin-top: 2%">
          <input type="text" class="form-control" v-model="searchParams"/>
          <button v-on:click="searchProfiles" class="btn-info btn-round" style="padding: 2.5%; border: solid #01131f 1pt">Search</button>
        </div>
      </li>
    </template>
  </navbar>
</template>

<script>
import { DropDown, Navbar, NavLink } from '@/components';
import { Popover } from 'element-ui';
import router from "../router";
export default {
  name: 'main-navbar',
  props: {
    transparent: Boolean,
    colorOnScroll: Number
  },
  components: {
    DropDown,
    Navbar,
    NavLink,
    [Popover.name]: Popover
  },
  data(){
    return {
      user: {},
      searchParams: ''
    }
  },
  mounted: function() {
    this.setupData()
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
      router.push('/')
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
};
</script>

<style scoped></style>
