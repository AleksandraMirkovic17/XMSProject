import Vue from 'vue'
import { BootstrapVue, BootstrapVueIcons } from 'bootstrap-vue'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(BootstrapVue)
Vue.use(BootstrapVueIcons)
import App from './App.vue'
import router from './router'
import NowUiKit from './plugins/now-ui-kit';

Vue.config.productionTip = false
Vue.use(NowUiKit)

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')






