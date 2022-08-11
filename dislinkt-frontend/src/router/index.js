import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Registration from "../views/Registration";
import Home from "../views/Home";
import ProfileSearchView from "../views/ProfileSearchView";
import Profile from "../views/Profile";
import Connections from "../views/Connections";
import Profile2 from "../views/Profile2";
import MainNavbar from "../layout/MainNavbar";
import MainFooter from "../layout/MainFooter";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/home',
    name: 'Home',
    components: {default: Home, header: MainNavbar, footer: MainFooter},
    props: {
      header: {colorOnScroll: 400},
      footer: {backgroundColor: 'black'}
    }
  },
  {
    path: '/login',
    name: 'Login',
    components: {default: Login, header: MainNavbar},
    props: {
      header: {colorOnScroll: 400},
    }
  },
  {
    path: '/register',
    name: 'Register',
    components: {default: Registration, header: MainNavbar, footer: MainFooter},
    props: {
      header: {colorOnScroll: 400},
      footer: {backgroundColor: 'black'}
    }
  },
  {
    path: '/profiles/:search',
    name: 'ProfileView',
    components: {default: ProfileSearchView, header: MainNavbar, footer: MainFooter}
  },
  {
    path: '/profile/:id',
    name: 'Profile',
    components: {default: Profile, header: MainNavbar, footer: MainFooter}
  },
  {
    path: '/connections',
    name: 'Connections',
    components: {default: Connections, header: MainNavbar, footer: MainFooter},
    props: {
      header: {colorOnScroll: 400},
      footer: {backgroundColor: 'black'}
    }
  },
  {
    path: '/profile2',
    name: 'Profile2',
    component: Profile2
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
  scrollBehavior: to => {
  if (to.hash) {
    return { selector: to.hash };
  } else {
    return { x: 0, y: 0 };
  }
},
})

export default router
