import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../components/HomePage.vue'
import UserRegistration from '../components/UserRegistration.vue'
import UserLogin from '../components/UserLogin.vue'

const routes = [
    {
        path: '/',
        name: 'home',
        component: HomePage
    },
      {
        path: '/home',
        name: 'HomePage',
        component: HomePage
      },
      {
        path: '/userRegistration',
        name: 'userRegistration',
        component: UserRegistration
       },
       {
        path: '/userLogin',
        name: 'userLogin',
        component: UserLogin
       }
      ]


      const router = createRouter({
        history: createWebHistory(process.env.BASE_URL),
        routes
      })
      
      
      
      export default router