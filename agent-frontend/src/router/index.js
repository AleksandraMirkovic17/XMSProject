import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../components/HomePage.vue'
import UserRegistration from '../components/UserRegistration.vue'
import UserLogin from '../components/UserLogin.vue'
import RegistrationCompany from '../components/RegistrationCompany.vue'
import ConfirmCompany from '../components/ConfirmCompany.vue'
import UserHomePage from '../components/UserHomePage.vue'
import ListingCompany from '../components/ListingCompany.vue'
import OwnerCompany from '../components/OwnerCompany.vue'



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
       },
       {
        path: '/registrationCompany',
        name: 'registrationCompany',
        component: RegistrationCompany
       },
       {
        path: '/confirmCompany',
        name: 'confirmCompany',
        component: ConfirmCompany
       },

       {
        path: '/listingCompany',
        name: 'listingCompany',
        component: ListingCompany
       },

       {
        path: '/userHomePage',
        name: 'userHomePage',
        component: UserHomePage
       },
       {
        path: '/ownerCompany',
        name: 'ownerCompany',
        component: OwnerCompany
       }



      ]


      const router = createRouter({
        history: createWebHistory(process.env.BASE_URL),
        routes
      })
      
      
      
      export default router