import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      alias: ['/index.html'],
      // component: HomeView
      component: () => import('@/views/CategoryView.vue')
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/categories',
      name: 'categories',
      component: () => import('@/views/CategoryView.vue')
    },
    {
      path: '/categories/create',
      name: 'categories.create',
      component: () => import('@/views/CategoryForm.vue')
    },
    {
      path: '/wallets',
      name: 'wallets',
      component: () => import('@/views/WalletView.vue')
    },
    {
      path: '/wallet/create',
      name: 'wallets.create',
      component: () => import('@/views/WalletForm.vue')
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('@/views/NotFoundView.vue')
    }
  ]
})

export default router
