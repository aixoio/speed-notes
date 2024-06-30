import { createRouter, createWebHashHistory, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import("../views/HomeView.vue"),
    },
    {
      path: '/signup',
      name: 'signup',
      component: () => import("../views/SignupView.vue"),
    },
    {
      path: '/login',
      name: 'login',
      component: () => import("../views/LoginView.vue"),
    },
    {
      path: '/dash',
      name: 'dash',
      component: () => import("../views/dash/DashView.vue"),
    },
    {
      path: '/newnote',
      name: 'newnote',
      component: () => import("../views/dash/NewNoteView.vue"),
    },
  ]
})

export default router
