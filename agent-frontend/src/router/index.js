import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../components/HomePage.vue'
import UserRegistration from '../components/UserRegistration.vue'

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
       }]


      const router = createRouter({
        history: createWebHistory(process.env.BASE_URL),
        routes
      })
      
      
      
      export default router