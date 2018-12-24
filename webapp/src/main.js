import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'

import iView from 'iview'
import 'iview/dist/styles/iview.css'
import store from './store'
Vue.use(iView)

Vue.prototype.$http = axios
Vue.config.productionTip = false

new Vue({
	router,
	store,
	render: h => h(App)
}).$mount('#app')