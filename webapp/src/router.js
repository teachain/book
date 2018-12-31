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
    path: '/device/add',
    name: 'device',
    component: () =>
      import ('./views/AddDevice.vue')
  }, {
    path: '/driver/add',
    name: 'driver',
    component: () =>
      import ('./views/AddDriver.vue')
  }, {
    path: '/model/add',
    name: 'modelAdd',
    component: () =>
      import ('./views/AddModel.vue')
  }, {
    path: '/record/add',
    name: 'record',
    component: () =>
      import ('./views/AddRecord.vue')
  }, {
    path: '/record/list',
    name: 'record',
    component: () =>
      import ('./views/ListRecord.vue')
  }, {
    path: '/maintain/add',
    name: 'maintainadd',
    component: () =>
      import ('./views/AddMaintain.vue')
  }, {
    path: '/maintain/list',
    name: 'maintainlist',
    component: () =>
      import ('./views/ListMaintain.vue')
  }, {
    path: '/gas/add',
    name: 'gasadd',
    component: () =>
      import ('./views/AddGas.vue')
  }, {
    path: '/gas/list',
    name: 'gaslist',
    component: () =>
      import ('./views/ListGas.vue')
  }]
})