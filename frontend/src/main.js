import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue/dist/bootstrap-vue.esm'
import App from './App.vue'
import router from './router'
import  Axios  from  'axios'
import VueCookie from 'vue-cookie'
Vue.use(VueCookie)

// set default config
//VueCookies.config('7d')
//import VueMeta from 'vue-meta'

Vue.config.productionTip = false
Vue.prototype.$http  =  Axios;
const  token  =   VueCookie.get('token');

//VueCookies.set('token','1123123');

if (token) {
  Vue.prototype.$http.defaults.headers.common['Authorization'] = token
}

// Import the styles directly. (Or you could add them via script tags.)
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

Vue.use(BootstrapVue);

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
