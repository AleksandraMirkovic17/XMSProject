import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Registration from "../views/Registration";
import Home from "../views/Home";
import ProfileSearchView from "../views/ProfileSearchView";
import Profile from "../views/Profile";
import Connections from "../views/Connections";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/home',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Registration
  },
  {
    path: '/profiles/:search',
    name: 'ProfileView',
    component: ProfileSearchView
  },
  {
    path: '/profile/:id',
    name: 'Profile',
    component: Profile
  },
  {
    path: '/connections',
    name: 'Connections',
    component: Connections
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
