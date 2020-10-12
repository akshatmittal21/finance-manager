import Vue from 'vue'
import VueRouter from 'vue-router'
import FinanceManager from '@/views/FinanceManager.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'FinanceManager',

    component: FinanceManager,
    meta: {
      title: 'FinanceManager',
    },
  },
  {
    path: '/add-new-account',
    name: 'AddNewAccount',
    // route level code-splitting
    // this generates a separate chunk (AddNewAccount.[hash].js) for this route
    // which is lazy-loaded when the route is visited.

    meta: {
      title: 'Add New Account',
    },
    component: () =>
      import(
        /* webpackChunkName: "AddNewAccount" */ '../views/AddNewAccount.vue'
      ),
  },
  {
    path: '/edit-account/:accountId',
    name: 'EditAccount',
    // route level code-splitting
    // this generates a separate chunk (EditAccount.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    meta: {
      title: 'Edit Account',
    },
    component: () =>
      import(/* webpackChunkName: "EditAccount" */ '../views/EditAccount.vue'),
  },
  {
    path: '/add-category',
    name: 'AddCategory',
    // route level code-splitting
    // this generates a separate chunk (AddCategory.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    meta: {
      title: 'Add Category',
    },
    component: () =>
      import(/* webpackChunkName: "AddCategory" */ '../views/AddCategory.vue'),
  },
  {
    path: '/edit-category/:categoryId',
    name: 'EditCategory',
    // route level code-splitting
    // this generates a separate chunk (EditCategory.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    meta: {
      title: 'Edit Category',
    },
    component: () =>
      import(
        /* webpackChunkName: "EditCategory" */ '../views/EditCategory.vue'
      ),
  },
]

const router = new VueRouter({
  routes,
})

export default router
