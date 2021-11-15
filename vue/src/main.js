import Vue from 'vue'

import 'normalize.css/normalize.css' // A modern alternative to CSS resets

import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import locale from 'element-ui/lib/locale/lang/en'

import VueEasyLightbox from 'vue-easy-lightbox'

import '@/styles/index.scss' // global css

import App from './App'
import store from './store'
import router from './router'

import '@/icons' // icon
import '@/permission' // permission control

import LoadScript from 'vue-plugin-load-script'

import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'

// Import Bootstrap an BootstrapVue CSS files (order is important)
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'


// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue)
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin)
Vue.use(LoadScript)

Vue.use(ElementUI, {locale})
Vue.use(VueEasyLightbox)


Vue.loadScript("/assets/static/bootstrap/js/bootstrap.min.js")
    .then(() => {
      // Script is loaded, do something
    })
    .catch(() => {
      // Failed to fetch script
    })

Vue.loadScript("/assets/static/js/chart.min.js")
.then(() => {
  // Script is loaded, do something
})
.catch(() => {
  // Failed to fetch script
});

Vue.loadScript("/assets/static/js/bs-init.js")
.then(() => {
  // Script is loaded, do something
})
.catch(() => {
  // Failed to fetch script
});

Vue.loadScript("/assets/static/js/theme.js")
.then(() => {
  // Script is loaded, do something
})
.catch(() => {
  // Failed to fetch script
});

Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
