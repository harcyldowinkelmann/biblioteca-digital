import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/entrar',
    name: 'entrar',
    component: () => import("@/views/Login.vue")
  },
  {
    path: '/index',
    name: 'index',
    component: () => import("@/views/IndexView.vue")
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: '/home',
    name: 'home',
    component: () => import("@/views/HomeView.vue")
  },
  {
    path: '/cadastro',
    name: 'cadastro',
    component: () => import("@/views/CadastroView.vue")
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
