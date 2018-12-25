import Vue from 'vue'
import Router from 'vue-router'
//import Home from './views/Home.vue'

Vue.use(Router)

export default new Router({
  //mode: 'history',
  base: process.env.BASE_URL,
  routes: [{
    path: '/',
    name: 'home',
    component: () =>
      import ('./views/Home.vue')
  }, {
    path: '/about',
    name: 'about',
    component: () =>
      import ('./views/About.vue')
  }, {
    path: '/device',
    name: 'device',
    component: () =>
      import ('./views/Device.vue')
  }, {
    path: '/driver',
    name: 'driver',
    component: () =>
      import ('./views/Driver.vue')
  }, {
    path: '/model',
    name: 'model',
    component: () =>
      import ('./views/Model.vue')
  },{
    path: '/record',
    name: 'record',
    component: () =>
      import ('./views/Record.vue')
  }]
})